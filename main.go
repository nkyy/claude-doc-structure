package main

import (
	"os"

	"github.com/claude-code/claude-doc-structure/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}