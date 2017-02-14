package server // import "github.com/Zenika/goru/server"

import (
	"net/http"
	"path"
	"strconv"

	"github.com/pkg/errors"

	"github.com/husobee/vestigo"
	"github.com/spf13/viper"
)

func StartRouter(port int) error {
	router := vestigo.NewRouter()

	contextPath := viper.GetString("server.contextPath")
	if !path.IsAbs(contextPath) {
		return errors.New("Context path must an absolute path")
	}

	router.Post(path.Join(contextPath, "document/:file/edit"), postEditeurHandler)
	router.Put(path.Join(contextPath, "document/:file/content"), putDocumentHandler)
	router.Get(path.Join(contextPath, "document/:file/content"), getDocumentHandler)

	println("Starting server on port", port, "and context path", contextPath)

	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		return errors.Wrap(err, "Could not start HTTP server")
	}

	return nil
}
