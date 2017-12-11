package handlers

import (
	"net/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type handlerError struct {
	errors []string
}

func handleError(w http.ResponseWriter, err error) {
	if err == nil {
		// this should not happen.
		log.Panic("nil error passed")
	}

	// log the error.
	// TODO: Add the url to the log.
	log.Errorf("error while processing request: %s", err.Error())

	// Send a 500 error.
	w.WriteHeader(http.StatusInternalServerError)
	err2 := writeError(w, err)

	if err2 != nil {
		// there was an error writing the response.
		log.Errorf("Exception while writing error: %s", err2.Error())
		log.Panic(err)
	}
}

func writeError(w http.ResponseWriter, err error) error {
	if err == nil {
		log.Panic("nil error passed")
	}
	encoder := json.NewEncoder(w)
	e := handlerError{
		errors: []string{err.Error()},
	}
	return encoder.Encode(e)
}
