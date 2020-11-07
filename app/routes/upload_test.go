package routes_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"github.com/nikolalohinski/gofsud/app/configuration"
	"github.com/nikolalohinski/gofsud/app/routes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Upload", func() {
	const (
		directoryToUse   = "/tmp"
		filepathToUse    = "folder/file.ext"
		fullFilePath     = directoryToUse + "/" + filepathToUse
		urlToUse         = "http://127.0.0.1:8080/api/vX/files/" + filepathToUse
		fileContentToUse = "I should be uploaded"
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

		request = MustReturn(http.NewRequest(http.MethodPut, urlToUse, strings.NewReader(fileContentToUse))).(*http.Request)
		request = mux.SetURLVars(request, map[string]string{routes.FilePathKey: filepathToUse})

		handler = routes.NewHandler(configuration.Configuration{
			Directory: directoryToUse,
		})
	})

	JustBeforeEach(func() {
		handler.Upload(responseRecorder, request)
	})

	AfterEach(func() {
		_ = os.RemoveAll(fullFilePath)
	})

	Context("nominal case", func() {
		It("should upload the file correctly", func() {
			Expect(responseRecorder.Code).To(Equal(http.StatusCreated)) // language=JSON
			Expect(responseRecorder.Body).To(MatchJSON(`{
				"message": "file folder/file.ext successfully uploaded",
				"status": 201
			}`))
			writtenContent := MustReturn(ioutil.ReadFile(fullFilePath)).([]byte)
			Expect(string(writtenContent)).To(Equal(fileContentToUse))
		})
	})

	Context("when path is invalid", func() {
		BeforeEach(func() {
			request = mux.SetURLVars(request, nil)
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
})
