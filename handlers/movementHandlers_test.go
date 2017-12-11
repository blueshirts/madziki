package handlers

import (
	"testing"
	"net/http"
	"github.com/blueshirts/madziki/api"
	"encoding/json"
	"bytes"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

const url = "http://localhost:3000/movements"
const name = "TestName"
const description = "TestDescription"
const details = "TestDetails"
const user = "TestUser"

// The ID of the created movement.
var ID bson.ObjectId

func init() {
	log.SetLevel(log.DebugLevel)
}

func TestPostMovementHandler(t *testing.T) {
	movement := api.Movement{
		Name:        name,
		Description: description,
		Details:     details,
		User:        user,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(movement)

	t.Logf("posting new movement to url: %s", url)
	res, err := http.Post(url, "application/json; charset=utf-8", b)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Status code: %d", res.StatusCode)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Bad response code: %d", res.StatusCode)
	}

	var result api.Movement
	readMovement(res.Body, &result)
	if result.ID.Hex() == "" {
		t.Fatalf("found empty movement.ID value")
	} else {
		t.Logf("POST'd movement with ID: %s", result.ID.Hex())
	}

	ID = result.ID

	if result.Name != name {
		t.Fatalf("incorrect value for movement.Name: %s", result.Name)
	}
	if result.Description != description {
		t.Fatalf("incorrect value for movement.Description: %s", result.Description)
	}
	if result.Details != details {
		t.Fatalf("incorrect value for movement.Details: %s", result.Details)
	}
	if result.User != user {
		t.Fatalf("incorrect value for movement.User: %s", result.User)
	}
}

func TestPostMovementHandler_NoName(t *testing.T) {
	movement := api.Movement{}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(movement)

	res, err := http.Post(url, "application/json; charset=utf-8", b)
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("expected status code to be 500: %d", res.StatusCode)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMovementHandler(t *testing.T) {
	t.Logf("Current ID: %s", ID.Hex())
	getUrl := fmt.Sprintf("%s/%s", url, ID.Hex())
	t.Logf("Requesting url: %s", getUrl)
	r, err := http.Get(getUrl)
	if err != nil {
		t.Fatal(err)
	}
	if r.StatusCode != http.StatusOK {
		t.Fatalf("Invalid http status code: %d", r.StatusCode)
	}
	var m api.Movement
	err = readMovement(r.Body, &m)

	defer deleteMovement(t)

	if err != nil {
		t.Fatal(err)
	}
	log.Debug(m)
}

func deleteMovement(t *testing.T) {
	if ID.Hex() != "" {
		t.Logf("deleting movement with id: %s", ID.Hex())
		err := api.DeleteMovement(ID)
		if err != nil {
			t.Fatalf("error while deleting movement: %s", err.Error())
		}
	}
}
