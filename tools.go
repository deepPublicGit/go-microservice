//go:build tools
// +build tools

package go_microservice

import (
	// https://golangci-lint.run
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	// https://go.dev/blog/vuln
	_ "golang.org/x/vuln/cmd/govulncheck"
)
