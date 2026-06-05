package services

import (
	"context"
	"crypto/ecdh"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Access-Grid/accessgrid-go/client"
	"github.com/Access-Grid/accessgrid-go/models"
)

// Captured wire-compat fixture — same envelope + caller keypair used in
// Ruby / Elixir / JS / Python / Java / PHP specs. caller_private_key is
// ephemeral and single-use by design (the server rejects reuse on pubkey
// fingerprint), so committing it carries no credential risk.
const fixtureCallerPrivateKeyPEM = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIou+Kk08kWAjhi0WyIx+L2GrgStGBCPODlwKYKd5BydoAoGCCqGSM49
AwEHoUQDQgAE+gnDxXJt1SBaCK8roKH8QvOa/ItdQUe85JIsUc6RvhD/udLaFtHY
m+MnOmeSdVaKTPWudH0+iGbleB3kS7lYxQ==
-----END EC PRIVATE KEY-----
`

const fixtureExpectedPlaintext = "FIXTURE-PLAINTEXT-NOT-A-CREDENTIAL"

func fixtureEnvelope() map[string]interface{} {
	return map[string]interface{}{
		"alg":                  "ECDH-ES+A256GCM",
		"ephemeral_public_key": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7mg6i99GcIVutMPr/PXSBSQVlbLM\ntnJO10ZBjk9ZTfw6wwAVNBnDBiqY7VrdOG1JdFOYoac+NkAlyMRGYk2tVQ==\n-----END PUBLIC KEY-----\n",
		"iv":                   "5X2OCht+kLB/xQmX",
		"ciphertext":           "ckYyA3FdRYjOFI/FKz/QeR5Yf9nZZFzo73kDXKZSB/EgbQ==",
		"tag":                  "0vwkjVaCwi5zl37xvJPxeg==",
	}
}

func loadFixturePrivateKey(t *testing.T) *ecdh.PrivateKey {
	t.Helper()
	block, _ := pem.Decode([]byte(fixtureCallerPrivateKeyPEM))
	if block == nil {
		t.Fatal("failed to decode fixture private key PEM")
	}
	ecdsaPriv, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("ParseECPrivateKey: %v", err)
	}
	priv, err := ecdsaPriv.ECDH()
	if err != nil {
		t.Fatalf("ECDH conversion: %v", err)
	}
	return priv
}

func TestDecryptEnvelope_DecryptsCapturedFixture(t *testing.T) {
	priv := loadFixturePrivateKey(t)
	plaintext, err := decryptEnvelope(fixtureEnvelope(), priv)
	if err != nil {
		t.Fatalf("decryptEnvelope: %v", err)
	}
	if string(plaintext) != fixtureExpectedPlaintext {
		t.Errorf("plaintext = %q, want %q", string(plaintext), fixtureExpectedPlaintext)
	}
}

func TestDecryptEnvelope_TamperedTag_ReturnsErrDecrypt(t *testing.T) {
	priv := loadFixturePrivateKey(t)
	env := fixtureEnvelope()
	rawTag, _ := base64.StdEncoding.DecodeString(env["tag"].(string))
	rawTag[0] ^= 0x01
	env["tag"] = base64.StdEncoding.EncodeToString(rawTag)

	_, err := decryptEnvelope(env, priv)
	if !errors.Is(err, models.ErrDecryptFailed) {
		t.Errorf("err = %v, want errors.Is(err, ErrDecryptFailed)", err)
	}
}

func TestDecryptEnvelope_WrongPrivKey_ReturnsErrDecrypt(t *testing.T) {
	wrong, _, err := generateKeypair()
	if err != nil {
		t.Fatalf("generateKeypair: %v", err)
	}

	_, err = decryptEnvelope(fixtureEnvelope(), wrong)
	if !errors.Is(err, models.ErrDecryptFailed) {
		t.Errorf("err = %v, want errors.Is(err, ErrDecryptFailed)", err)
	}
}

func TestDecryptEnvelope_MissingEphemeralPubKey_ReturnsErrInvalidEnvelope(t *testing.T) {
	priv := loadFixturePrivateKey(t)
	env := fixtureEnvelope()
	delete(env, "ephemeral_public_key")

	_, err := decryptEnvelope(env, priv)
	if !errors.Is(err, models.ErrInvalidEnvelope) {
		t.Errorf("err = %v, want errors.Is(err, ErrInvalidEnvelope)", err)
	}
}

func TestDecryptEnvelope_NonBase64IV_ReturnsErrInvalidEnvelope(t *testing.T) {
	priv := loadFixturePrivateKey(t)
	env := fixtureEnvelope()
	env["iv"] = "not!base64!"

	_, err := decryptEnvelope(env, priv)
	if !errors.Is(err, models.ErrInvalidEnvelope) {
		t.Errorf("err = %v, want errors.Is(err, ErrInvalidEnvelope)", err)
	}
}

func TestGenerateKeypair_ReturnsDistinctKeypairs(t *testing.T) {
	_, pemA, err := generateKeypair()
	if err != nil {
		t.Fatalf("generateKeypair: %v", err)
	}
	_, pemB, err := generateKeypair()
	if err != nil {
		t.Fatalf("generateKeypair: %v", err)
	}
	if pemA == pemB {
		t.Error("two generateKeypair calls returned identical public PEMs")
	}
}

func TestConsoleService_RevealSmartTap_DecryptsServerEnvelope(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/console/card-templates/tmpl-42/smart-tap/reveal" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"key_version": "tmpl-42",
			"collector_id": "12345678",
			"fingerprint": "sha256:deadbeef",
			"encrypted_private_key": {
				"alg": "ECDH-ES+A256GCM",
				"ephemeral_public_key": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7mg6i99GcIVutMPr/PXSBSQVlbLM\ntnJO10ZBjk9ZTfw6wwAVNBnDBiqY7VrdOG1JdFOYoac+NkAlyMRGYk2tVQ==\n-----END PUBLIC KEY-----\n",
				"iv": "5X2OCht+kLB/xQmX",
				"ciphertext": "ckYyA3FdRYjOFI/FKz/QeR5Yf9nZZFzo73kDXKZSB/EgbQ==",
				"tag": "0vwkjVaCwi5zl37xvJPxeg=="
			}
		}`))
	}))
	defer server.Close()

	c, err := client.NewClient("acct", "secret", client.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	console := NewConsoleService(c)
	priv := loadFixturePrivateKey(t)

	result, err := console.revealSmartTapWithKey(context.Background(), "tmpl-42", priv)
	if err != nil {
		t.Fatalf("revealSmartTapWithKey: %v", err)
	}
	if result.KeyVersion != "tmpl-42" {
		t.Errorf("KeyVersion = %q, want tmpl-42", result.KeyVersion)
	}
	if result.CollectorID != "12345678" {
		t.Errorf("CollectorID = %q, want 12345678", result.CollectorID)
	}
	if result.Fingerprint != "sha256:deadbeef" {
		t.Errorf("Fingerprint = %q, want sha256:deadbeef", result.Fingerprint)
	}
	if result.PrivateKey != fixtureExpectedPlaintext {
		t.Errorf("PrivateKey = %q, want %q", result.PrivateKey, fixtureExpectedPlaintext)
	}
}

func TestConsoleService_PublishTemplate_HappyPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/console/card-templates/tmpl-42/publish" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"tmpl-42","status":"in-review"}`))
	}))
	defer server.Close()

	c, err := client.NewClient("acct", "secret", client.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	console := NewConsoleService(c)

	result, err := console.PublishTemplate(context.Background(), "tmpl-42")
	if err != nil {
		t.Fatalf("PublishTemplate: %v", err)
	}
	if result.ID != "tmpl-42" {
		t.Errorf("ID = %q, want tmpl-42", result.ID)
	}
	if result.Status != "in-review" {
		t.Errorf("Status = %q, want in-review", result.Status)
	}
}
