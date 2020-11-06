package routes

import (
	"encoding/json"
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

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(map[string]interface{}{
		"message": err.Error(),
		"status":  status,
	}); err != nil {
		log.Error(errors.Wrap(err, "failed to write error body"))
	}
}

func writeSuccess(w http.ResponseWriter, message string, status int) {
	log.Info(message)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(map[string]interface{}{
		"message": message,
		"status":  status,
	}); err != nil {
		log.Error(errors.Wrap(err, "failed to write success body"))
	}
}
