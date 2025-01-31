package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (h *fileHandler) Delete(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)[FilePathKey]
	if !ok {
		writeError(w, errors.New("missing file path"), http.StatusBadRequest)

		return
	}

	log.Infof("About to delete %s file", name)

	path := filepath.Clean(filepath.Join(h.Configuration.Directory, name))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if os.IsNotExist(err) {
			writeError(w, errors.Wrap(err, "file does not exist"), http.StatusNotFound)

			return
		}

		writeError(w, errors.Wrap(err, "failed to stat file"), http.StatusInternalServerError)

		return
	}

	if err := os.Remove(path); err != nil {
		writeError(w, errors.Wrap(err, "failed to delete file"), http.StatusInternalServerError)

		return
	}

	writeSuccess(w, fmt.Sprintf("file %s successfully deleted", name), http.StatusOK)
}
