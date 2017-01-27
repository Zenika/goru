package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/husobee/vestigo"

	"github.com/Zenika/pdf-api/domain"
	"github.com/Zenika/pdf-api/pdf"
)

func postEditeurHandler(w http.ResponseWriter, r *http.Request) {
	if err := handlePostEditeurRequest(r); err != nil {
		log.Println(err)
		//FIXME write error to response
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

	ts := time.Now().UnixNano()
	outputFile := fmt.Sprintf("%d.pdf", ts)

	if err := pdf.ApplyActionsToFile(file, actions, outputFile); err != nil {
		return err
	}

	return nil
}
