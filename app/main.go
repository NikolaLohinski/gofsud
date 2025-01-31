package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/nikolalohinski/gofsud/app/configuration"
	"github.com/nikolalohinski/gofsud/app/routes"
	log "github.com/sirupsen/logrus"
)

var (
	config      configuration.Configuration
	formatter   = new(log.JSONFormatter)
	fileHandler routes.FileHandler
	server      = new(http.Server)
)

func initConfiguration() {
	var err error
	config, err = configuration.LoadConfiguration()
	checkErrorAndExit(err)
}

func initLog() {
	formatter.PrettyPrint = config.LogPrettyPrint
	formatter.FieldMap = log.FieldMap{
		log.FieldKeyMsg:  "message",
		log.FieldKeyFunc: "version",
		log.FieldKeyFile: "name",
	}
	formatter.CallerPrettyfier = func(_ *runtime.Frame) (function string, file string) {
		return config.ServiceVersion, config.ServiceName
	}
	log.SetFormatter(formatter)
	log.SetLevel(config.LogLevel)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
}

func initServer() {
	server.Addr = fmt.Sprintf(":%v", config.ServicePort)
}

func initHandler() {
	fileHandler = routes.NewHandler(config)
}

func main() {
	initConfiguration()
	initLog()
	initServer()
	initHandler()

	router := mux.NewRouter()
	subRouter := router.PathPrefix(fmt.Sprintf("/api/v%v", config.GetAPIVersion())).Subrouter()

	subRouter.HandleFunc("/files/{"+routes.FilePathKey+":.*}", fileHandler.Download).Methods(http.MethodGet)
	subRouter.HandleFunc("/files/{"+routes.FilePathKey+":.*}", fileHandler.Delete).Methods(http.MethodDelete)
	subRouter.HandleFunc("/files/{"+routes.FilePathKey+":.*}", fileHandler.Upload).Methods(http.MethodPut)

	server.Handler = router

	log.Infof("starting server on port %v", config.ServicePort)
	checkErrorAndExit(server.ListenAndServe())
}

func checkErrorAndExit(err error) {
	if err == nil {
		return
	}

	log.Fatal(err.Error())
}
