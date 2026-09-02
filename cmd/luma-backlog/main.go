// Command luma-backlog is the entry point. It holds no logic: everything
// lives under internal/, because the contract is the command line rather
// than a Go package (docs/spec.md §9a.1).
package main

import (
	"os"

	"github.com/lumastack/luma-backlog/internal/cli"
)

func main() {
	os.Exit(cli.Main(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
