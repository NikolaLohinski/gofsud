package routes

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (h *fileHandler) Download(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)[FilePathKey]
	if !ok {
		writeError(w, errors.New("missing file path"), http.StatusBadRequest)

		return
	}

	log.Infof("About to download %s file", name)

	path := filepath.Clean(filepath.Join(h.Configuration.Directory, name))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		writeError(w, errors.Wrap(err, "file does not exist"), http.StatusNotFound)

		return
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		writeError(w, errors.Wrap(err, "failed to open file"), http.StatusInternalServerError)

		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Errorf("failed to close file: %v", err.Error())
		}
	}()

	if _, err := io.Copy(w, file); err != nil {
		writeError(w, errors.Wrap(err, "failed to copy file content"), http.StatusInternalServerError)

		return
	}

	log.Infof("file %s successfully downloaded", name)
	w.WriteHeader(http.StatusOK)
}
