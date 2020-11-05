// +build mage

package main

import (
	"github.com/fatih/color"
	"github.com/magefile/mage/sh"
)

var Default = Build

func Build() {
	color.Red("# Installing tools")
	if err := sh.RunV("go", "run", "-mod=mod", "github.com/izumin5210/gex/cmd/gex", "--build"); err != nil {
		panic(err)
	}
}
