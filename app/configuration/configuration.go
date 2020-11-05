package configuration

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	EnvPrefix = "GOFSUD"
)

// Build information. Populated at build-time.
var (
	ServiceName    = "unknown"
	ServiceVersion = "X.X.X"
)

type Configuration struct {
	ServiceName    string
	ServiceVersion string
	ServicePort    int       `split_words:"true" default:"8080"`
	LogLevel       log.Level `split_words:"true" default:"DEBUG"`
	LogPrettyPrint bool      `split_words:"true" default:"true"`
	Directory      string    `split_words:"true" default:"/tmp"`
}

func (c Configuration) GetAPIVersion() string {
	return strings.Split(c.ServiceVersion, ".")[0]
}

func LoadConfiguration() (configuration Configuration, err error) {
	if err := envconfig.Process(EnvPrefix, &configuration); err != nil {
		return configuration, errors.Wrap(err, "failed to parse env vars")
	}
	if configuration.ServiceName == "" {
		configuration.ServiceName = ServiceName
	}
	if configuration.ServiceVersion == "" {
		configuration.ServiceVersion = ServiceVersion
	}

	return configuration, nil
}
