package main // import "github.com/Zenika/pdf-api"

import (
	"os"
	"strconv"

	"github.com/pkg/errors"

	"github.com/urfave/cli"

	"github.com/Zenika/pdf-api/domain"
	"github.com/Zenika/pdf-api/pdf"
	"github.com/Zenika/pdf-api/server"
)

func main() {
	app := cli.NewApp()

	app.Name = "pdf-api"
	app.Usage = "Manipulate PDF files"

	app.Commands = []cli.Command{
		{
			Name:      "left-rotate-page",
			Usage:     "Rotate a page anticlockwise",
			ArgsUsage: "<input-file> <page-number> <output-file>",
			Action: func(c *cli.Context) error {
				inputFile := c.Args().Get(0)
				pageNumber, err := strconv.Atoi(c.Args().Get(1))
				if err != nil {
					return errors.Wrap(err, "Page number must be a valid integer")
				}
				outputFile := c.Args().Get(2)
				action := domain.Action{
					Action: "LEFT_ROTATE_PAGE",
					Page:   pageNumber,
				}
				return applyActionToFile(inputFile, action, outputFile)
			},
		},
		{
			Name:      "right-rotate-page",
			Usage:     "Rotate a page clockwise",
			ArgsUsage: "<input-file> <page-number> <output-file>",
			Action: func(c *cli.Context) error {
				inputFile := c.Args().Get(0)
				pageNumber, err := strconv.Atoi(c.Args().Get(1))
				if err != nil {
					return errors.Wrap(err, "Page number must be a valid integer")
				}
				outputFile := c.Args().Get(2)
				action := domain.Action{
					Action: "RIGHT_ROTATE_PAGE",
					Page:   pageNumber,
				}
				return applyActionToFile(inputFile, action, outputFile)
			},
		},
		{
			Name:      "delete-page",
			Usage:     "Delete a page",
			ArgsUsage: "<input-file> <page-number> <output-file>",
			Action: func(c *cli.Context) error {
				inputFile := c.Args().Get(0)
				pageNumber, err := strconv.Atoi(c.Args().Get(1))
				if err != nil {
					return errors.Wrap(err, "Page number must be a valid integer")
				}
				outputFile := c.Args().Get(2)
				action := domain.Action{
					Action: "DELETE_PAGE",
					Page:   pageNumber,
				}
				return applyActionToFile(inputFile, action, outputFile)
			},
		},
		{
			Name:      "move-page",
			Usage:     "Move a page",
			ArgsUsage: "<input-file> <page-number> <target> <output-file>",
			Action: func(c *cli.Context) error {
				inputFile := c.Args().Get(0)
				pageNumber, err := strconv.Atoi(c.Args().Get(1))
				if err != nil {
					return errors.Wrap(err, "Page number must be a valid integer")
				}
				target, err := strconv.Atoi(c.Args().Get(2))
				if err != nil {
					return errors.Wrap(err, "Target must be a valid integer")
				}
				outputFile := c.Args().Get(3)
				action := domain.Action{
					Action: "MOVE_PAGE",
					Page:   pageNumber,
					Target: &target,
				}
				return applyActionToFile(inputFile, action, outputFile)
			},
		},
		{
			Name:      "server",
			Usage:     "Start the server",
			ArgsUsage: "<port>",
			Action: func(c *cli.Context) error {
				port, err := strconv.Atoi(c.Args().Get(0))
				if err != nil {
					return errors.Wrap(err, "Port must be a valid integer")
				}
				if err = pdf.EnsureDocumentsDir(); err != nil {
					return err
				}
				return server.StartRouter(port)
			},
		},
	}

	app.EnableBashCompletion = true

	app.Run(os.Args)
}

func applyActionToFile(inputFile string, action domain.Action, outputFile string) error {
	actions := make([]domain.Action, 1)
	actions[0] = action

	return pdf.ApplyActionsToFile(inputFile, actions, outputFile)
}
