package pdf // import "github.com/Zenika/goru/pdf"

import (
	"github.com/pkg/errors"
	"github.com/unidoc/unidoc/pdf"
)

type Document struct {
	Pages []*pdf.PdfPage
}

func (document *Document) isValidPage(page int) error {
	if page < 0 || page >= len(document.Pages) {
		return errors.Errorf("Invalid page number %d, should be between 1 and %d", page+1, len(document.Pages))
	}
	return nil
}
