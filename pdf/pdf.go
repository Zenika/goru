package pdf // import "github.com/Zenika/goru/pdf"

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/unidoc/unidoc/pdf"

	"github.com/Zenika/goru/domain"
)

type Document struct {
	Pages []*pdf.PdfPage
}

//FIXME split in several files

func EnsureDocumentsDir() error {
	if err := os.MkdirAll(getDocumentsPath(), 0755); err != nil {
		return errors.Wrap(err, "Error while creating documents dir")
	}

	return nil
}

func ApplyActionsToFile(inputFile string, actions []domain.Action, outputFile string) error {
	document, err := readDocumentFromFile(inputFile)
	if err != nil {
		return err
	}

	if err = document.applyActionsToDocument(actions); err != nil {
		return err
	}

	if err = document.writeDocumentToFile(outputFile); err != nil {
		return err
	}

	return nil
}

func ApplyActionToFile(inputFile string, action domain.Action, outputFile string) error {
	actions := make([]domain.Action, 1)
	actions[0] = action

	return ApplyActionsToFile(inputFile, actions, outputFile)
}

func GetDocumentPath(document string) string {
	return filepath.Join(getDocumentsPath(), document+".pdf")
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

func (document *Document) applyActionsToDocument(actions []domain.Action) error {
	for _, action := range actions {
		page := action.Page - 1

		//FIXME replace by a map
		switch action.Action {
		case "LEFT_ROTATE_PAGE":
			if err := document.rotatePage(page, -90); err != nil {
				return err
			}
		case "RIGHT_ROTATE_PAGE":
			if err := document.rotatePage(page, 90); err != nil {
				return err
			}
			break
		case "DELETE_PAGE":
			if err := document.deletePage(page); err != nil {
				return err
			}
			break
		case "MOVE_PAGE":
			if err := document.movePage(page, action.Target); err != nil {
				return err
			}
			break
		default:
			return errors.Errorf("Unknown action %s", action.Action)
		}
	}

	return nil
}

func (document *Document) rotatePage(page int, rotation int) error {
	if err := document.isValidPage(page); err != nil {
		return err
	}

	var pageRotation int64 = 0
	if document.Pages[page].Rotate != nil {
		pageRotation = *document.Pages[page].Rotate
	}
	pageRotation += int64(rotation)
	document.Pages[page].Rotate = &pageRotation

	return nil
}

func (document *Document) deletePage(page int) error {
	if err := document.isValidPage(page); err != nil {
		return err
	}

	document.Pages = append(document.Pages[:page], document.Pages[page+1:]...)

	return nil
}

func (document *Document) movePage(page int, pTarget *int) error {
	if err := document.isValidPage(page); err != nil {
		return err
	}

	if pTarget == nil {
		return errors.New("Missing mandatory target parameter for move page action")
	}

	target := *pTarget - 1

	if target < 0 || target > len(document.Pages) {
		return errors.Errorf("Invalid target %d, should be between 1 and %d", target+1, len(document.Pages)+1)
	}

	if page < target {
		target -= 1
	}

	pageToMove := document.Pages[page]
	document.Pages = append(document.Pages[:page], document.Pages[page+1:]...)
	document.Pages = append(document.Pages[:target], append([]*pdf.PdfPage{pageToMove}, document.Pages[target:]...)...)

	return nil
}

func (document *Document) isValidPage(page int) error {
	if page < 0 || page >= len(document.Pages) {
		return errors.Errorf("Invalid page number %d, should be between 1 and %d", page+1, len(document.Pages))
	}
	return nil
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

func getDocumentsPath() string {
	return viper.GetString("server.documentsPath")
}
