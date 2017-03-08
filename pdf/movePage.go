package pdf // import "github.com/Zenika/goru/pdf"

import (
	"github.com/Zenika/goru/domain"
	"github.com/pkg/errors"
	"github.com/unidoc/unidoc/pdf"
)

const MOVE_PAGE = "MOVE_PAGE"

func movePage(document *Document, action domain.Action) error {
	page := action.GetPage()

	if err := document.isValidPage(page); err != nil {
		return err
	}

	if action.Target == nil {
		return errors.New("Missing mandatory target parameter for move page action")
	}

	target := *action.Target - 1

	if target < 0 || target > len(document.Pages) {
		return errors.Errorf("Invalid target %d, should be between 1 and %d", target+1, len(document.Pages)+1)
	}

	pageToMove := document.Pages[page]
	document.Pages = append(document.Pages[:page], document.Pages[page+1:]...)
	document.Pages = append(document.Pages[:target], append([]*pdf.PdfPage{pageToMove}, document.Pages[target:]...)...)

	return nil
}

func init() {
	actionFuncs[MOVE_PAGE] = movePage
}
