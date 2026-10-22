package release

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateVersion(t *testing.T) {
	for _, version := range []string{"1.5.5", "0.1.0"} {
		if err := ValidateVersion(version); err != nil {
			t.Errorf("ValidateVersion(%q): %v", version, err)
		}
	}
	for _, version := range []string{"v1.5.5", "1.5", "01.5.5"} {
		if err := ValidateVersion(version); err == nil {
			t.Errorf("ValidateVersion(%q) accepted invalid input", version)
		}
	}
}

func TestValidateRepository(t *testing.T) {
	root := t.TempDir()
	write := func(name, value string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(root, name), []byte(value), 0600); err != nil {
			t.Fatal(err)
		}
	}
	write("VERSION", "1.5.5\n")
	write("RELEASE_NOTES.md", "## 1.5.5 - 2026-10-22\n\n- Updated behavior.\n")
	write("BlessingBuddy.toc", "## Version: 1.5.5\n")
	if got, err := ValidateRepository(root); err != nil || got != "1.5.5" {
		t.Fatalf("ValidateRepository() = %q, %v", got, err)
	}
	write("BlessingBuddy.toc", "## Version: 1.5.4\n")
	if _, err := ValidateRepository(root); err == nil {
		t.Fatal("ValidateRepository accepted a mismatched TOC version")
	}
}
