package routes_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/nikolalohinski/gofsud/app/configuration"
	"github.com/nikolalohinski/gofsud/app/routes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete", func() {
	const (
		directoryToUse   = "/tmp"
		filepathToUse    = "folder/file.ext"
		fullFilePath     = directoryToUse + "/" + filepathToUse
		urlToUse         = "http://127.0.0.1:8080/api/vX/files/" + filepathToUse
		fileContentToUse = "I should be deleted"
	)

	var (
		request          *http.Request
		responseRecorder *httptest.ResponseRecorder

		handler routes.FileHandler
	)
	BeforeEach(func() {
		responseRecorder = &httptest.ResponseRecorder{
			Body: bytes.NewBuffer(nil),
		}

		request = MustReturn(http.NewRequest(http.MethodDelete, urlToUse, nil)).(*http.Request)
		request = mux.SetURLVars(request, map[string]string{routes.FilePathKey: filepathToUse})

		handler = routes.NewHandler(configuration.Configuration{
			Directory: directoryToUse,
		})

		Must(os.MkdirAll(filepath.Dir(fullFilePath), os.ModePerm))
		file := MustReturn(os.Create(fullFilePath)).(*os.File)
		MustReturn(file.WriteString(fileContentToUse))
		Must(file.Close())
	})

	JustBeforeEach(func() {
		handler.Delete(responseRecorder, request)
	})

	Context("nominal case", func() {
		It("should delete the file correctly", func() {
			Expect(responseRecorder.Code).To(Equal(http.StatusOK))
			// language=JSON
			Expect(responseRecorder.Body).To(MatchJSON(`{
				"message": "file folder/file.ext successfully deleted",
				"status": 200
			}`))
		})
	})

	Context("when path is invalid", func() {
		BeforeEach(func() {
			request = mux.SetURLVars(request, nil)
		})

		AfterEach(func() {
			Must(os.RemoveAll(fullFilePath))
		})

		It("should return an error", func() {
			Expect(responseRecorder.Code).To(Equal(http.StatusBadRequest))
			// language=JSON
			Expect(responseRecorder.Body).To(MatchJSON(`{
				"message": "missing file path",
				"status": 400
			}`))
		})
	})

	Context("when file does not exist", func() {
		BeforeEach(func() {
			Must(os.RemoveAll(fullFilePath))
		})

		It("should return an error", func() {
			Expect(responseRecorder.Code).To(Equal(http.StatusNotFound))
			// language=JSON
			Expect(responseRecorder.Body).To(MatchJSON(`{
				"message": "file does not exist: stat /tmp/folder/file.ext: no such file or directory",
				"status": 404
			}`))
		})
	})
})
