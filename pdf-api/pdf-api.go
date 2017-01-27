package main

import (
	"os"
	"strconv"

	"github.com/pkg/errors"

	"github.com/unidoc/unidoc/pdf"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "pdf-api"
	app.Usage = "Manipulate PDF files"

	app.Commands = []cli.Command{
		{
			Name:      "rotate-page",
			Usage:     "Rotate a page",
			ArgsUsage: "<input-file> <page-number> <rotation> <output-file>",
			Action: func(c *cli.Context) error {
				inputFile := c.Args().Get(0)
				pageNumber, err := strconv.Atoi(c.Args().Get(1))
				if err != nil {
					return errors.Wrap(err, "Page number must be a valid integer")
				}
				rotation, err := strconv.Atoi(c.Args().Get(2))
				if err != nil {
					return errors.Wrap(err, "Rotation number must be a valid integer")
				}
				outputFile := c.Args().Get(3)
				return rotatePage(inputFile, pageNumber, rotation, outputFile)
			},
		},
	}
	app.EnableBashCompletion = true

	app.Run(os.Args)
}

func rotatePage(inputFile string, pageNumber int, rotation int, outputFile string) error {
	in, err := os.Open(inputFile)
	if err != nil {
		return errors.Wrap(err, "Error while opening input PDF file")
	}
	defer in.Close()

	reader, err := pdf.NewPdfReader(in)
	if err != nil {
		return errors.Wrap(err, "Error while creating PDF reader")
	}

	isEncrypted, err := reader.IsEncrypted()
	if err != nil {
		return errors.Wrap(err, "Error while determining if input PDF file is encrypted")
	}
	if isEncrypted {
		return errors.New("Input PDF file is encrypted")
	}

	numPages, err := reader.GetNumPages()
	if err != nil {
		return errors.Wrap(err, "Error while determining number of pages of input PDF file")
	}
	if pageNumber > numPages {
		return errors.Errorf("Invalid page number %s, must be <= to number of pages %s", pageNumber, numPages)
	}

	writer := pdf.NewPdfWriter()

	for curPageNumber := 1; curPageNumber <= numPages; curPageNumber++ {
		page, err := reader.GetPageAsPdfPage(curPageNumber)
		if err != nil {
			return errors.Wrap(err, "Error while reading page from input PDF file")
		}

		if curPageNumber == pageNumber {
			var pageRotation int64 = 0
			if page.Rotate != nil {
				pageRotation = *(page.Rotate)
			}
			pageRotation += int64(rotation)
			page.Rotate = &pageRotation
		}

		err = writer.AddPage(page.GetPageAsIndirectObject())
		if err != nil {
			return errors.Wrap(err, "Error while writing page to output PDF file")
		}
	}

	out, err := os.Create(outputFile)
	if err != nil {
		return errors.Wrap(err, "Error while opening output PDF file")
	}
	defer out.Close()

	err = writer.Write(out)
	if err != nil {
		return errors.Wrap(err, "Error while writing output PDF file")
	}

	return nil
}
