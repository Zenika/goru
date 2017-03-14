package pdf // import "github.com/Zenika/goru/pdf"

import (
	"github.com/Zenika/goru/domain"
)

const (
	LEFT_ROTATE_PAGE  = "LEFT_ROTATE_PAGE"
	RIGHT_ROTATE_PAGE = "RIGHT_ROTATE_PAGE"
)

type rotatePageFunc func(*Document, domain.Action, int) error

func (rotatePage rotatePageFunc) apply(rotation int) actionFunc {
	return func(document *Document, action domain.Action) error {
		return rotatePage(document, action, rotation)
	}
}

var rotatePage rotatePageFunc = func(document *Document, action domain.Action, rotation int) error {
	page := action.GetPage()

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

func init() {
	actionFuncs[LEFT_ROTATE_PAGE] = rotatePage.apply(-90)
	actionFuncs[RIGHT_ROTATE_PAGE] = rotatePage.apply(90)
}
