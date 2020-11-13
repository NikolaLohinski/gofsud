package spellbook

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"
)

type Go mg.Namespace

// Run formatting tools on code base
func (Go) Format() error {
	color.Cyan("# Formatting everything...")

	args := []string{"-s", "-w"}
	args = append(args, getGoFiles()...)

	return sh.RunV("gofumpt", args...)
}

// Run linter on code base
func (Go) Lint() error {
	mg.Deps(Go.Format)
	color.Cyan("# Linting code...")

	return sh.RunV("golangci-lint", "run")
}

// Clean imports
func (Go) Tidy() error {
	fmt.Println("# Cleaning go modules...")

	return sh.RunV("go", "mod", "tidy", "-v")
}

// Vendor dependencies
func (Go) Dependencies() error {
	color.Cyan("# Vendoring dependencies...")

	return sh.RunV("go", "mod", "vendor")
}

// Run unit tests
func (Go) Test() error {
	color.Cyan("# Running unit tests...")

	if err := sh.Run("mkdir", "-p", ".local/junit"); err != nil {
		return errors.Wrap(err, "failed to create junit folder")
	}

	return sh.RunV("gotestsum", "--junitfile", ".local/unit-tests.xml", "--", "-short", "-cover", "./app/...")
}

func getGoFiles() []string {
	var goFiles []string

	err := filepath.Walk(os.Getenv("PWD"), func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "vendor/") {
			return filepath.SkipDir
		}
		if strings.Contains(path, "tools/") {
			return filepath.SkipDir
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		absPath := strings.Replace(path, os.Getenv("PWD"), ".", 1)
		goFiles = append(goFiles, absPath)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return goFiles
}
