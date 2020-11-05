package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (h *fileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)[FilePathKey]
	if !ok {
		writeError(w, errors.New("missing file path"), http.StatusBadRequest)
		return
	}

	path := filepath.Clean(filepath.Join(h.Configuration.Directory, name))
	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		writeError(w, errors.Wrap(err, "failed to create directory tree"), http.StatusInternalServerError)
		return
	}

	log.Infof("file %s successfully uploaded", name)
}
