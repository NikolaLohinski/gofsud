package spellbook

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Bin mg.Namespace

func (Bin) Build() error {
	color.Cyan("# Building app...")

	version := determineVersion()

	varsSetByLinker := map[string]string{
		"github.com/nikolalohinski/gofsud/app/configuration.ServiceVersion": version,
		"github.com/nikolalohinski/gofsud/app/configuration.ServiceName":    "gofsud",
	}

	linkerArgs := make([]string, len(varsSetByLinker))
	for name, value := range varsSetByLinker {
		linkerArgs = append(linkerArgs, "-X", fmt.Sprintf("%s=%s", name, value))
	}

	linkerArgs = append(linkerArgs, "-s", "-w")

	return sh.RunWith(map[string]string{
		"CGO_ENABLED": "0",
	}, "go", "build", "-ldflags", strings.Join(linkerArgs, " "), "-mod=vendor", "-o", ".local/bin/gofsud", "./app")
}

func (b Bin) Run() error {
	if err := b.Build(); err != nil {
		return err
	}

	color.Cyan("# Starting app...")

	return sh.RunV(".local/bin/gofsud")
}
