package cmd // import "github.com/Zenika/goru/cmd"

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Zenika/goru/domain"
	"github.com/Zenika/goru/pdf"
)

var rightRotatePageCmd = &cobra.Command{
	Use:   "right-rotate-page <input-file> <page-number> <output-file>",
	Short: "Rotate a page clockwise",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("right-rotate-page needs 3 arguments")
		}
		inputFile := args[0]
		pageNumber, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.Wrap(err, "Page number must be a valid integer")
		}
		outputFile := args[2]
		action := domain.NewAction(pdf.RIGHT_ROTATE_PAGE, pageNumber, nil)
		return pdf.ApplyActionToFile(inputFile, action, outputFile)
	},
}

func init() {
	RootCmd.AddCommand(rightRotatePageCmd)
}
