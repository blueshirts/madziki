package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/blueshirts/madziki/api"
	log "github.com/sirupsen/logrus"
	"io"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// writeMovement writes a movement instance to the response.
func writeMovement(w http.ResponseWriter, movement api.Movement) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(movement)
}

// readMovement reads a movement instance from the body.
func readMovement(body io.ReadCloser, movement *api.Movement) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(&movement)
}

func PostMovementHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("received post movement request...")

	var movement api.Movement
	err := readMovement(r.Body, &movement)
	if err != nil {
		handleError(w, err)
	}
	log.Debug(fmt.Sprintf("Movement: %+v", movement))
	id, err := api.CreateMovement(movement)
	log.Debugf("Created new movement with id: %s", id.Hex())
	if err != nil {
		handleError(w, err)
		return
	} else {
		movement.ID = id
		err = writeMovement(w, movement)
	}
}

func GetMovementHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		// TODO: Return 400 error.
	}

	log.Debugf("Retrieving movement for id: %s", id)
	var m api.Movement
	err := api.GetMovement(bson.ObjectIdHex(id), &m)
	if err != nil {
		handleError(w, err)
		return
	}
	err = writeMovement(w, m)
	if err != nil {
		handleError(w, err)
		return
	}
}
