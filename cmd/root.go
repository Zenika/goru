package cmd // import "github.com/Zenika/goru/cmd"

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/goru/config"
	"github.com/Zenika/goru/log"
)

var RootCmd = &cobra.Command{
	Use:           "goru",
	Short:         "Goru is a simple PDF manipulation tool",
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := config.Init(); err != nil {
			return err
		}
		if err := log.Init(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&config.File, "config", "", "Config file")
	RootCmd.PersistentFlags().String("logLevel", "INFO", "Log level")

	viper.BindPFlag("logLevel", RootCmd.PersistentFlags().Lookup("logLevel"))
}
