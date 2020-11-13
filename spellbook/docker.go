package spellbook

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Docker mg.Namespace

// Build docker image.
func (Docker) Package() error {
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
	return sh.RunV("docker", "push", "${IMAGE_DESTINATION:-theagentk/gofsud}")
}
