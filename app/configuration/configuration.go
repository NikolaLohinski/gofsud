package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Configuration struct{}

func LoadConfiguration() (configuration Configuration, err error) {
	for _, a := range os.Args {
		if a != "--help" {
			continue
		}

		if err := envconfig.Usage("GOFSUD", &configuration); err != nil {
			return configuration, errors.Wrap(err, "failed to parse usage")
		}
		os.Exit(0)
	}

	if err := envconfig.Process("GOFSUD", &configuration); err != nil {
		return configuration, errors.Wrap(err, "failed to parse env vars")
	}

	return configuration, nil
}
