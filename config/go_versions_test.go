package config

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// TARGET VERSION - Update this when upgrading Go
// ══════════════════════════════════════════════════════════════════════════════
const TARGET_GO = "1.23.5"

func rootDir(t *testing.T) string {
	t.Helper()
	// config/ is one level below the repo root
	dir, err := filepath.Abs(filepath.Join("..", "."))
	if err != nil {
		t.Fatalf("failed to resolve root dir: %v", err)
	}
	return dir
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read %s: %v", path, err)
	}
	return string(data)
}

func majorMinor(version string) string {
	parts := strings.SplitN(version, ".", 3)
	if len(parts) < 2 {
		return version
	}
	return parts[0] + "." + parts[1]
}

func TestToolVersionsMatchesTarget(t *testing.T) {
	root := rootDir(t)
	content := readFile(t, filepath.Join(root, ".tool-versions"))

	var goVersion string
	for _, line := range strings.Split(content, "\n") {
		parts := strings.Fields(line)
		if len(parts) == 2 && parts[0] == "golang" {
			goVersion = parts[1]
			break
		}
	}

	if goVersion == "" {
		t.Fatal(".tool-versions does not contain a golang entry")
	}

	if goVersion != TARGET_GO {
		t.Errorf(".tool-versions golang = %q, want %q", goVersion, TARGET_GO)
	}
}

func TestGoModMatchesTarget(t *testing.T) {
	root := rootDir(t)
	content := readFile(t, filepath.Join(root, "go.mod"))

	re := regexp.MustCompile(`(?m)^go\s+(\S+)`)
	match := re.FindStringSubmatch(content)
	if match == nil {
		t.Fatal("go.mod does not contain a go directive")
	}

	goModVersion := match[1]
	targetMM := majorMinor(TARGET_GO)

	if goModVersion != targetMM && goModVersion != TARGET_GO {
		t.Errorf("go.mod go directive = %q, want %q or %q", goModVersion, targetMM, TARGET_GO)
	}
}

func TestCIMatrixIncludesTarget(t *testing.T) {
	root := rootDir(t)
	content := readFile(t, filepath.Join(root, ".github", "workflows", "ci.yml"))

	// Extract versions from go: ['1.23', '1.24', '1.25'] matrix
	re := regexp.MustCompile(`go:\s*\[([^\]]+)\]`)
	match := re.FindStringSubmatch(content)
	if match == nil {
		t.Fatal("CI workflow does not contain a go version matrix")
	}

	raw := match[1]
	var versions []string
	for _, v := range strings.Split(raw, ",") {
		v = strings.TrimSpace(v)
		v = strings.Trim(v, "'\"")
		if v != "" {
			versions = append(versions, v)
		}
	}

	targetMM := majorMinor(TARGET_GO)
	found := false
	for _, v := range versions {
		if v == targetMM {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("CI matrix %v does not include target major.minor %q", versions, targetMM)
	}
}
