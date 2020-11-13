// +build mage

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	// mage:import
	"github.com/nikolalohinski/gofsud/spellbook"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Verify

func init() {
	// Add local bin in PATH
	tooling()
}

func tooling() error {
	name, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	p, err := filepath.Abs(path.Join(name, "tools", "bin"))
	if err != nil {
		panic(err)
	}

	return os.Setenv("PATH", fmt.Sprintf("%s:%s", p, os.Getenv("PATH")))
}

// Get and install required tools for development.
func Install() error {
	return sh.RunV("mage", "-d", "./tools")
}

// Validate code base.
func Verify() {
	mg.SerialDeps(
		spellbook.Go.Tidy,
		spellbook.Go.Dependencies,
		spellbook.Go.Lint,
		spellbook.Go.Test,
	)
}
