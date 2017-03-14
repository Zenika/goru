package pdf // import "github.com/Zenika/goru/pdf"

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/unidoc/unidoc/pdf"
)

var (
	documentsPath *string
)

func EnsureDocumentsDir() error {
	if err := os.MkdirAll(getDocumentsPath(), 0755); err != nil {
		return errors.Wrap(err, "Error while creating documents dir")
	}

	return nil
}

func GetDocumentPath(document string) string {
	return filepath.Join(getDocumentsPath(), document+".pdf")
}

func getDocumentsPath() string {
	if documentsPath == nil {
		documentsPathValue := viper.GetString("server.documentsPath")
		documentsPath = &documentsPathValue
	}
	return *documentsPath
}

func readDocumentFromFile(file string) (*Document, error) {
	in, err := os.Open(file)
	if err != nil {
		return nil, errors.Wrap(err, "Error while opening input PDF file")
	}
	defer in.Close()

	reader, err := pdf.NewPdfReader(in)
	if err != nil {
		return nil, errors.Wrap(err, "Error while creating PDF reader")
	}

	isEncrypted, err := reader.IsEncrypted()
	if err != nil {
		return nil, errors.Wrap(err, "Error while determining if input PDF file is encrypted")
	}
	if isEncrypted {
		return nil, errors.New("Input PDF file is encrypted")
	}

	numPages, err := reader.GetNumPages()
	if err != nil {
		return nil, errors.Wrap(err, "Error while determining number of pages of input PDF file")
	}

	document := Document{}

	for curPageNumber := 1; curPageNumber <= numPages; curPageNumber++ {
		page, err := reader.GetPageAsPdfPage(curPageNumber)
		if err != nil {
			return nil, errors.Wrap(err, "Error while reading page from input PDF file")
		}

		document.Pages = append(document.Pages, page)
	}

	return &document, nil
}

func (document *Document) writeDocumentToFile(file string) error {
	numPages := len(document.Pages)

	writer := pdf.NewPdfWriter()

	for curPageNumber := 0; curPageNumber < numPages; curPageNumber++ {
		page := document.Pages[curPageNumber]

		if err := writer.AddPage(page.GetPageAsIndirectObject()); err != nil {
			return errors.Wrap(err, "Error while writing page to output PDF file")
		}
	}

	out, err := os.Create(file)
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
