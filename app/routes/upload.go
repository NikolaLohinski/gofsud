package routes

import (
	"fmt"
	"io"
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

	log.Infof("About to upload %s file", name)

	path := filepath.Clean(filepath.Join(h.Configuration.Directory, name))
	directory := filepath.Dir(path)

	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		writeError(w, errors.Wrap(err, "failed to create directory tree"), http.StatusInternalServerError)

		return
	}

	file, err := os.Create(path)
	if err != nil {
		writeError(w, errors.Wrap(err, "failed to create file"), http.StatusInternalServerError)

		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Error(errors.Wrap(err, "failed to close request body"))
		}
	}()

	if _, err := io.Copy(file, r.Body); err != nil {
		writeError(w, errors.Wrap(err, "failed to copy body content into file"), http.StatusInternalServerError)

		return
	}

	log.Infof("file %s successfully uploaded", name)

	writeSuccess(w, fmt.Sprintf("file %s successfully uploaded", name), http.StatusCreated)
}
