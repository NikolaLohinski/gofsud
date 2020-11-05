package main

import (
	"fmt"
	"log"

	"github.com/nikolalohinski/gofsud/app/configuration"
)

var (
	config configuration.Configuration
)

func init() {
	var err error
	config, err = configuration.LoadConfiguration()
	checkErrorAndExit(err)
}

func checkErrorAndExit(err error) {
	if err == nil {
		return
	}
	log.Fatal(err.Error())
}

func main() {
	fmt.Println("Hello world!", config)
}
