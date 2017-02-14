package cmd // import "github.com/Zenika/goru/cmd"

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goru",
	Short: "Goru is a simple PDF manipulation tool",
}
