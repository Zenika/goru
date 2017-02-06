package server // import "github.com/Zenika/goru/server"

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/husobee/vestigo"
)

func StartRouter(port int) error {
	router := vestigo.NewRouter()

	router.Post("/api/document/:file/edit", postEditeurHandler)
	router.Put("/api/document/:file/content", putDocumentHandler)
	router.Get("/api/document/:file/content", getDocumentHandler)

	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		return errors.Wrap(err, "Could not start HTTP server")
	}

	return nil
}
