package main

import (
	"github.com/jenkins-x-labs/jx-labs/cmd/root"
	"os"
)

// Entrypoint for the command
func main() {
	if err := root.Run(nil); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
