package domain // import "github.com/Zenika/goru/domain"

type Action struct {
	Action string
	Page   int
	Target *int
}

func NewAction(action string, page int, target *int) Action {
	return Action{
		Action: action,
		Page:   page,
		Target: target,
	}
}

func (action *Action) GetPage() int {
	return action.Page - 1
}
