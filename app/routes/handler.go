package routes

import (
	"fmt"
	"net/http"

	"github.com/nikolalohinski/gofsud/app/configuration"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	FilePathKey = "path"
)

func NewHandler(config configuration.Configuration) FileHandler {
	return &fileHandler{
		Configuration: config,
	}
}

type FileHandler interface {
	Delete(w http.ResponseWriter, r *http.Request)
	Upload(w http.ResponseWriter, r *http.Request)
	Download(w http.ResponseWriter, r *http.Request)
}

type fileHandler struct {
	Configuration configuration.Configuration
}

func writeError(w http.ResponseWriter, err error, status int) {
	log.Error(err)

	w.WriteHeader(status)

	if _, err := w.Write([]byte(fmt.Sprintf("%v Status %s\n", status, http.StatusText(status)))); err != nil {
		log.Error(errors.Wrap(err, "failed to write error body"))
	}
}
