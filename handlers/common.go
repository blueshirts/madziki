package handlers

import (
	"net/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)

type handlerError struct {
	Errors []string
}

func sendError(w http.ResponseWriter, code int, err error) {
	// log the error.
	// TODO: Add the url to the log.
	log.Errorf("error while processing request: %s", err)

	// Send a 500 error.
	w.WriteHeader(code)
	err2 := writeError(w, err)

	if err2 != nil {
		// there was an error writing the response.
		log.Errorf("Exception while writing error: %s", err2.Error())
		log.Panic(err)
	}
}

func send500(w http.ResponseWriter, err error) {
	sendError(w, http.StatusInternalServerError, err)
}

func send404(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNotFound, err)
}

func send400(w http.ResponseWriter, err error) {
	sendError(w, http.StatusBadRequest, err)
}

func writeError(w http.ResponseWriter, err error) error {
	var e handlerError

	if _, ok := err.(validator.ValidationErrors); ok {
		if len(err.(validator.ValidationErrors)) == 0 {
			log.Panic("empty validation errors passed to writeError")
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()

			e.Errors = append(e.Errors, fmt.Sprintf("%s is %s", err.Field(), err.Kind()))
		}
	} else {
		// there was a single error
		e.Errors = []string{err.Error()}
	}
	encoder := json.NewEncoder(w)
	return encoder.Encode(e)
}
