package services

// Internal crypto helpers for the SmartTap reveal flow.
//
// Driven by ConsoleService.RevealSmartTap; not part of the public SDK
// surface. Pure stdlib — crypto/ecdh + crypto/hkdf + crypto/aes + crypto/cipher.

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/hkdf"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/Access-Grid/accessgrid-go/models"
)

const hkdfInfo = "accessgrid-smart-tap-reveal-v1"

// generateKeypair generates a fresh ephemeral P-256 keypair for a reveal
// call. Returns the private key (kept for decryptEnvelope) and the public
// key as a SubjectPublicKeyInfo PEM string (ready to submit).
func generateKeypair() (*ecdh.PrivateKey, string, error) {
	priv, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		return nil, "", fmt.Errorf("generate P-256 keypair: %w", err)
	}
	der, err := x509.MarshalPKIXPublicKey(priv.PublicKey())
	if err != nil {
		return nil, "", fmt.Errorf("marshal public key: %w", err)
	}
	pubPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: der})
	return priv, string(pubPEM), nil
}

// decryptEnvelope decrypts the encrypted_private_key envelope from the
// reveal endpoint. Returns the plaintext SmartTap PEM as bytes.
//
// Returns models.ErrInvalidEnvelope on missing/bad envelope fields, and
// models.ErrDecryptFailed on AES-GCM auth-tag verification failure.
func decryptEnvelope(envelope map[string]interface{}, priv *ecdh.PrivateKey) ([]byte, error) {
	pemStr, ok := envelope["ephemeral_public_key"].(string)
	if !ok || pemStr == "" {
		return nil, fmt.Errorf("%w: ephemeral_public_key missing or not a string", models.ErrInvalidEnvelope)
	}
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, fmt.Errorf("%w: ephemeral_public_key is not a PEM block", models.ErrInvalidEnvelope)
	}
	parsedPub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrInvalidEnvelope, err)
	}
	// ParsePKIXPublicKey returns *ecdsa.PublicKey for NIST curves like P-256;
	// convert to *ecdh.PublicKey for the ECDH agreement.
	var serverPub *ecdh.PublicKey
	switch pub := parsedPub.(type) {
	case *ecdh.PublicKey:
		serverPub = pub
	case *ecdsa.PublicKey:
		serverPub, err = pub.ECDH()
		if err != nil {
			return nil, fmt.Errorf("%w: ecdh conversion: %v", models.ErrInvalidEnvelope, err)
		}
	default:
		return nil, fmt.Errorf("%w: ephemeral_public_key is not an EC public key", models.ErrInvalidEnvelope)
	}

	iv, err := decodeB64(envelope, "iv")
	if err != nil {
		return nil, err
	}
	ciphertext, err := decodeB64(envelope, "ciphertext")
	if err != nil {
		return nil, err
	}
	tag, err := decodeB64(envelope, "tag")
	if err != nil {
		return nil, err
	}

	sharedSecret, err := priv.ECDH(serverPub)
	if err != nil {
		return nil, fmt.Errorf("%w: ECDH: %v", models.ErrInvalidEnvelope, err)
	}

	aesKey, err := hkdf.Key(sha256.New, sharedSecret, nil, hkdfInfo, 32)
	if err != nil {
		return nil, fmt.Errorf("HKDF: %w", err)
	}

	block2, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, fmt.Errorf("AES cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block2)
	if err != nil {
		return nil, fmt.Errorf("AES-GCM: %w", err)
	}

	plaintext, err := gcm.Open(nil, iv, append(ciphertext, tag...), nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrDecryptFailed, err)
	}
	return plaintext, nil
}

func decodeB64(envelope map[string]interface{}, key string) ([]byte, error) {
	raw, ok := envelope[key].(string)
	if !ok {
		return nil, fmt.Errorf("%w: envelope %s must be a base64 string", models.ErrInvalidEnvelope, key)
	}
	out, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, fmt.Errorf("%w: envelope %s must be base64-encoded", models.ErrInvalidEnvelope, key)
	}
	return out, nil
}
