package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/hlts2/ctxt/internal/cli"
)

//go:embed version.txt
var version string

func main() {
	if err := cli.Run(version, os.Args...); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred: %v", err)
		os.Exit(1)
	}
}
