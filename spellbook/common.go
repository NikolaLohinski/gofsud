package spellbook

import (
	"os"

	"github.com/magefile/mage/sh"
)

func determineVersion() string {
	version := os.Getenv("VCS_TAG")
	if version == "" {
		version, _ = sh.Output("git", "describe", "--tags")
	}

	if version == "" {
		version, _ = sh.Output("git", "rev-parse", "HEAD")
		version = "X.X.X-" + version
	}

	return version
}
