package cmd // import "github.com/Zenika/goru/cmd"

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Zenika/goru/pdf"
	"github.com/Zenika/goru/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := pdf.EnsureDocumentsDir(); err != nil {
			return err
		}
		port := viper.GetInt("server.port")
		return server.StartRouter(port)
	},
}

func init() {
	serverCmd.Flags().String("contextPath", "/goru", "Context path")
	serverCmd.Flags().String("documentsPath", "documents", "Path of directory containing documents")
	serverCmd.Flags().IntP("port", "p", 8080, "Listening port")

	RootCmd.AddCommand(serverCmd)

	viper.BindPFlag("server.contextPath", serverCmd.Flags().Lookup("contextPath"))
	viper.BindPFlag("server.documentsPath", serverCmd.Flags().Lookup("documentsPath"))
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
}
