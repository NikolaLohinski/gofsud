package routes_test

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)

	gopath := path.Join(
		filepath.ToSlash(os.Getenv("GOPATH")),
		"src",
	)

	_, file, _, _ := runtime.Caller(1)
	packageFullName := strings.TrimPrefix(path.Dir(file), gopath+"/")
	packageName := strings.Title(path.Base(packageFullName))

	RunSpecs(t, fmt.Sprintf("%v suite (%v)", packageName, packageFullName))
}

func MustReturn(i interface{}, err error) interface{} {
	Must(err)
	return i
}

func Must(err error) {
	if err == nil {
		return
	}
	panic(err)
}
