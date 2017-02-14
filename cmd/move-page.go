package cmd // import "github.com/Zenika/goru/cmd"

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/Zenika/goru/domain"
	"github.com/Zenika/goru/pdf"
)

var movePageCmd = &cobra.Command{
	Use:   "move-page <input-file> <page-number> <target> <output-file>",
	Short: "Move a page",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 4 {
			return errors.New("move-page needs 4 arguments")
		}
		inputFile := args[0]
		pageNumber, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.Wrap(err, "Page number must be a valid integer")
		}
		target, err := strconv.Atoi(args[2])
		if err != nil {
			return errors.Wrap(err, "Target must be a valid integer")
		}
		outputFile := args[3]
		action := domain.Action{
			Action: "MOVE_PAGE",
			Page:   pageNumber,
			Target: &target,
		}
		return pdf.ApplyActionToFile(inputFile, action, outputFile)
	},
}

func init() {
	RootCmd.AddCommand(movePageCmd)
}
