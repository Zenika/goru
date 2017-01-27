package server

import (
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/husobee/vestigo"
)

func StartRouter(port int) error {
	router := vestigo.NewRouter()

	router.Post("/document/:file/editeur", postEditeurHandler)

	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		return errors.Wrap(err, "Could not start HTTP server")
	}

	return nil
}
