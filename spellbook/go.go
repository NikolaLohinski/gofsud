package spellbook

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Go mg.Namespace

func (Go) Format() error {
	color.Cyan("# Formatting everything...")
	args := []string{"-s", "-w"}
	args = append(args, getGoFiles()...)
	return sh.RunV("gofumpt", args...)
}

func (Go) Lint() error {
	mg.Deps(Go.Format)
	color.Cyan("# Linting code...")
	return sh.RunV("golangci-lint", "run")
}

func (Go) Tidy() error {
	fmt.Println("# Cleaning go modules...")
	return sh.RunV("go", "mod", "tidy", "-v")
}

func (Go) Dependencies() error {
	color.Cyan("# Vendoring dependencies...")
	return sh.RunV("go", "mod", "vendor")
}

func (Go) Test() error {
	color.Cyan("# Running unit tests...")
	if err := sh.Run("mkdir", "-p", ".local/junit"); err != nil {
		return err
	}
	return sh.RunV("gotestsum", "--junitfile", ".local/unit-tests.xml", "--", "-short", "-cover", "./app/...")
}

func (Go) Build() error {
	color.Cyan("# Building app...")

	version := os.Getenv("VCS_TAG")
	if version == "" {
		version, _ = sh.Output("git", "describe", "--tags")
	}
	if version == "" {
		version, _ = sh.Output("git", "rev-parse", "HEAD")
		version = "X.X.X-" + version
	}

	varsSetByLinker := map[string]string{
		"github.com/nikolalohinski/gofsud/app/configuration.ServiceVersion": version,
		"github.com/nikolalohinski/gofsud/app/configuration.ServiceName":    "gofsud",
	}
	var linkerArgs []string
	for name, value := range varsSetByLinker {
		linkerArgs = append(linkerArgs, "-X", fmt.Sprintf("%s=%s", name, value))
	}
	linkerArgs = append(linkerArgs, "-s", "-w")

	return sh.RunWith(map[string]string{
		"CGO_ENABLED": "0",
	}, "go", "build", "-ldflags", strings.Join(linkerArgs, " "), "-mod=vendor", "-o", ".local/bin/gosfud", "./app")
}

func (g Go) Run() error {
	if err := g.Build(); err != nil {
		return err
	}
	color.Cyan("# Starting app...")
	return sh.RunV(".local/bin/gosfud")
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
