package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/blueshirts/madziki/api"
	log "github.com/sirupsen/logrus"
	"io"
	"github.com/gorilla/mux"
	"errors"
	"time"
)

type movementListResult struct {
	Size    int            `json:"size"`
	Offset  int            `json:"offset"`
	Results []api.Movement `json:"results"`
}

// writeData writes data to the response in JSON format.
func writeData(w http.ResponseWriter, data interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	return encoder.Encode(data)
}

// readMovement reads a movement instance from the body.
func readMovement(body io.ReadCloser, movement *api.Movement) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(movement)
}

func readData(body io.ReadCloser, data interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(data)
}

func PostMovementHandler(w http.ResponseWriter, r *http.Request) {
	var m api.Movement
	err := readMovement(r.Body, &m)
	if err != nil {
		log.Panic("error while reading parsing movement data", err)
	}

	// set defaults
	m.Created = time.Now()
	m.Updated = time.Now()

	err = api.ValidateMovement(m)
	if err != nil {
		send400(w, err)
		return
	}

	id, err := api.CreateMovement(m)
	if err != nil {
		log.Panic("error while creating movement", err)
	}
	m.ID = id
	err = writeData(w, m)
}

func GetMovementHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		send400(w, errors.New("id is required"))
		return
	}

	log.Debugf("Retrieving movement for id: %s", id)
	var m api.Movement
	err := api.GetMovement(id, &m)
	if err != nil {
		log.Panic(fmt.Sprintf("error while retrieving movement: %s", id), err)
	} else if m == (api.Movement{}) {
		// not found
		send404(w, errors.New(fmt.Sprintf("movement with id: %s was not found", id)))
		return
	}
	err = writeData(w, m)
	if err != nil {
		log.Panic(fmt.Sprintf("error while writing movement: %s", id), err)
	}
}

func UpdateMovementHandler(w http.ResponseWriter, r *http.Request) {
	var m api.Movement
	err := readMovement(r.Body, &m)
	if err != nil {
		log.Panic("error while reading m from request", err)
	}

	err = api.ValidateMovement(m)
	if err != nil {
		send400(w, err)
		return
	}

	err = api.UpdateMovement(m)
	if err != nil {
		log.Panic(fmt.Sprintf("error while updating movement: %s", m.ID))
	}
	err = writeData(w, m)
	if err != nil {
		send500(w, err)
		return
	}
}

func DeleteMovementHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		send400(w, errors.New("id is required"))
		return
	}
	err := api.DeleteMovement(id)
	if err != nil {
		log.Panic("error while deleting movement", err)
	}
	log.Info(fmt.Sprintf("deleted movement with id: %s", id))
}

// GetMovementsListHandler processes the list movements request.
func GetMovementsListHandler(w http.ResponseWriter, r *http.Request) {
	var a []api.Movement
	err := api.ListMovements(&a)
	if err != nil {
		log.Panic("error while retrieving movements", err)
	}

	if len(a) == 0 {
		// ensure the results is an empty array and not a null pointer.
		a = []api.Movement{}
	}
	out := movementListResult{
		Results: a,
		Size:    len(a),
		Offset:  0,
	}
	err = writeData(w, out)
	if err != nil {
		log.Panic(err)
	}
}
