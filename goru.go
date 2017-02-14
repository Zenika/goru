package main

import (
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/Zenika/goru/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Error(err)
		log.Error("Use \"goru --help\" for usage")
		os.Exit(-1)
	}
}
