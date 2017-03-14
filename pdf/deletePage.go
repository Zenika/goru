package pdf // import "github.com/Zenika/goru/pdf"

import (
	"github.com/Zenika/goru/domain"
)

const DELETE_PAGE = "DELETE_PAGE"

func deletePage(document *Document, action domain.Action) error {
	page := action.GetPage()

	if err := document.isValidPage(page); err != nil {
		return err
	}

	document.Pages = append(document.Pages[:page], document.Pages[page+1:]...)

	return nil
}

func init() {
	actionFuncs[DELETE_PAGE] = deletePage
}
