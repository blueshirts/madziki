package api

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
)

const movementName = "TestName"
const movementDescription = "TestDescription"
const movementDetails = "TestDetails"
const user = "TestUser"

// The id of the inserted movement.
var ID bson.ObjectId

func TestCreateMovement(t *testing.T) {
	movement := Movement{
		Name:        movementName,
		Description: movementDescription,
		Details:     movementDetails,
		User:        user,
	}
	var err error
	ID, err = CreateMovement(movement)
	if err != nil {
		t.Fatal(err)
	}
	if ID == "" {
		t.Fatal("ID is not defined")
	}
}

func TestCreateMovement_NoName(t *testing.T) {
	movement := Movement{}
	var err error
	_, err = CreateMovement(movement)
	if err == nil {
		t.Fatal("error should have been thrown when name is empty")
	}
}

func TestGetMovement(t *testing.T) {
	var movement Movement
	err := GetMovement(ID, &movement)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListMovements(t *testing.T) {
	var movements []Movement
	err := ListMovements(&movements)
	if err != nil {
		t.Fatal(err)
	}
	if len(movements) != 1 {
		t.Fatalf("expected only 1 movement in list: %d, %v", len(movements), movements)
	}
	if movements[0].Name != movementName {
		t.Fatalf("expected movement name to equal: %s but found %s", movementName, movements[0].Name)
	}
}

func TestDeleteMovement(t *testing.T) {
	err := DeleteMovement(ID)
	if err != nil {
		t.Fatal(err)
	}
}
