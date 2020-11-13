package spellbook

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Docker mg.Namespace

// Build docker image
func (Docker) Package() error {
	return sh.RunV(
		"docker",
		"build",
		".",
		"--build-arg", "GO_IMAGE_VERSION=${GO_IMAGE_VERSION:-1.15.3-alpine}",
		"--build-arg", "DISTROLESS_IMAGE=${DISTROLESS_IMAGE:-gcr.io/distroless/base-debian10}",
		"--build-arg", "DISTROLESS_VERSION=${DISTROLESS_VERSION:-nonroot}",
		"--tag", "${IMAGE_DESTINATION:-theagentk/gofsud}:"+determineVersion(),
		"--tag", "${IMAGE_DESTINATION:-theagentk/gofsud}:latest",
	)
}

// Upload docker image to a registry
func (Docker) Push() error {
	return sh.RunV("docker", "push", "${IMAGE_DESTINATION:-theagentk/gofsud}")
}
