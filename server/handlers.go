package server // import "github.com/Zenika/goru/server"

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/husobee/vestigo"
	"github.com/pkg/errors"

	"github.com/Zenika/goru/domain"
	"github.com/Zenika/goru/pdf"
)

func postEditeurHandler(w http.ResponseWriter, r *http.Request) {
	if err := handlePostEditeurRequest(r); err != nil {
		log.Error(err)
		//FIXME write error to response
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func handlePostEditeurRequest(r *http.Request) error {
	file := vestigo.Param(r, "file")

	decoder := json.NewDecoder(r.Body)
	var actions []domain.Action

	if err := decoder.Decode(&actions); err != nil {
		return errors.Wrap(err, "Invalid JSON in request body")
	}

	documentPath := pdf.GetDocumentPath(file)

	if err := pdf.ApplyActionsToFile(documentPath, actions, documentPath); err != nil {
		return err
	}

	return nil
}

func putDocumentHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	if contentType != "application/pdf" {
		w.WriteHeader(406)
		return
	}

	if err := handlePutDocument(r); err != nil {
		log.Error(err)
		//FIXME write error to response
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func handlePutDocument(r *http.Request) error {
	file := vestigo.Param(r, "file")

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "Error while reading file content from request body")
	}

	if err = ioutil.WriteFile(pdf.GetDocumentPath(file), content, 0644); err != nil {
		return errors.Wrap(err, "Error while writing file content to disk")
	}

	return nil
}

func getDocumentHandler(w http.ResponseWriter, r *http.Request) {
	file := vestigo.Param(r, "file")

	http.ServeFile(w, r, pdf.GetDocumentPath(file))
}
