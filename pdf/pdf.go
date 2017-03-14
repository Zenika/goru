package pdf // import "github.com/Zenika/goru/pdf"

import (
	"github.com/Zenika/goru/domain"
)

type actionFunc func(*Document, domain.Action) error

var actionFuncs = make(map[string]actionFunc)

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

func (document *Document) applyActionsToDocument(actions []domain.Action) error {
	for _, action := range actions {
		if err := actionFuncs[action.Action](document, action); err != nil {
			return err
		}
	}

	return nil
}
