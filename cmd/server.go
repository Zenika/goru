package cmd // import "github.com/Zenika/goru/cmd"

import (
	"github.com/pkg/errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/goru/pdf"
	"github.com/Zenika/goru/server"
)

var cfgFile string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initConfig(); err != nil {
			return err
		}
		if err := pdf.EnsureDocumentsDir(); err != nil {
			return err
		}
		port := viper.GetInt("server.port")
		return server.StartRouter(port)
	},
}

func init() {
	serverCmd.Flags().StringVar(&cfgFile, "config", "", "Config file")
	serverCmd.Flags().IntP("port", "p", 8080, "Listening port")
	serverCmd.Flags().String("documentsPath", "documents", "Path of directory containing documents")

	RootCmd.AddCommand(serverCmd)

	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.documentsPath", serverCmd.Flags().Lookup("documentsPath"))
}

func initConfig() error {
	viper.SetConfigName("goru")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, notFound := err.(viper.ConfigFileNotFoundError); notFound {
			return nil
		}
		return errors.Wrap(err, "Error while reading config file")
	}

	println("Using config file :", viper.ConfigFileUsed())

	return nil
}
