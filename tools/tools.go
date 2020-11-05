// +build tools

package tools

// tool dependencies
import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/onsi/ginkgo/ginkgo"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
	_ "mvdan.cc/gofumpt/gofumports"
)
