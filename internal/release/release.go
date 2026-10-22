package release

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var versionPattern = regexp.MustCompile(`^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$`)
var headingPattern = regexp.MustCompile(`(?m)^## (\d+\.\d+\.\d+)(?:\s+[-—].*)?$`)

func ValidateVersion(version string) error {
	if !versionPattern.MatchString(strings.TrimSpace(version)) {
		return fmt.Errorf("invalid version %q, expected major.minor.patch", version)
	}
	return nil
}

func ValidateNotes(root, version string) error {
	if err := ValidateVersion(version); err != nil {
		return err
	}
	content, err := os.ReadFile(filepath.Join(root, "RELEASE_NOTES.md"))
	if err != nil {
		return err
	}
	headings := headingPattern.FindAllStringSubmatch(string(content), -1)
	if len(headings) != 1 || headings[0][1] != strings.TrimSpace(version) {
		return fmt.Errorf("RELEASE_NOTES.md must contain exactly one heading for %s", version)
	}
	return nil
}

func ValidateRepository(root string) (string, error) {
	content, err := os.ReadFile(filepath.Join(root, "VERSION"))
	if err != nil {
		return "", err
	}
	version := strings.TrimSpace(string(content))
	if err := ValidateNotes(root, version); err != nil {
		return "", err
	}
	toc, err := os.ReadFile(filepath.Join(root, "BlessingBuddy.toc"))
	if err != nil {
		return "", err
	}
	tocVersion := regexp.MustCompile(`(?m)^## Version: (.+)$`).FindStringSubmatch(string(toc))
	if len(tocVersion) != 2 || tocVersion[1] != version {
		return "", fmt.Errorf("BlessingBuddy.toc version must match VERSION (%s)", version)
	}
	return version, nil
}
