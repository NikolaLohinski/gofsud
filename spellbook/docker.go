package spellbook

import (
	"os"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Docker mg.Namespace

// Build docker image.
func (Docker) Package() error {
	color.Cyan("# Packaging app...")

	goImageVersion := os.Getenv("GO_IMAGE_VERSION")
	if goImageVersion == "" {
		goImageVersion = "1.15.3-alpine"
	}

	distrolessImage := os.Getenv("DISTROLESS_IMAGE")
	if distrolessImage == "" {
		distrolessImage = "gcr.io/distroless/static"
	}

	distrolessVersion := os.Getenv("DISTROLESS_VERSION")
	if distrolessVersion == "" {
		distrolessVersion = "nonroot"
	}

	imageDestination := os.Getenv("IMAGE_DESTINATION")
	if imageDestination == "" {
		imageDestination = "theagentk/gofsud"
	}

	color.Cyan("Golang versions:    %s", goImageVersion)
	color.Cyan("Distroless image:   %s", distrolessImage)
	color.Cyan("Distroless version: %s", distrolessVersion)

	return sh.RunV(
		"docker",
		"build",
		".",
		"--build-arg", "GO_IMAGE_VERSION="+goImageVersion,
		"--build-arg", "DISTROLESS_IMAGE="+distrolessImage,
		"--build-arg", "DISTROLESS_VERSION="+distrolessVersion,
		"--tag", imageDestination+":"+determineVersion(),
		"--tag", imageDestination+":latest",
	)
}

// Upload docker image to a registry.
func (Docker) Push() error {
	color.Cyan("# Pushing image...")

	imageDestination := os.Getenv("IMAGE_DESTINATION")
	if imageDestination == "" {
		imageDestination = "theagentk/gofsud"
	}

	color.Cyan("Image destination:  %s", imageDestination)

	return sh.RunV("docker", "push", imageDestination)
}
