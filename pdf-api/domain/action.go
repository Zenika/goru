package domain // import "github.com/Zenika/pdf-api/domain"

type Action struct {
	Action string
	Page   int
	Target *int
}
