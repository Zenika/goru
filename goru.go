package main

import (
	"os"

	"github.com/Zenika/goru/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
