package cmd // import "github.com/Zenika/goru/cmd"

import (
	log "github.com/Sirupsen/logrus"
	"github.com/Zenika/goru/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Version: %s, Hash: %s", version.Version, version.Hash)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
