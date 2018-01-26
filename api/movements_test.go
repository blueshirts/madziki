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
	// make sure to use the module level ID
	ID, err = CreateMovement(movement)
	if err != nil {
		t.Fatal(err)
	}
	if ID == "" {
		t.Fatal("ID is not defined")
	}
	t.Logf("Created movement with id: %s", ID.Hex())
}

func TestCreateMovement_NoName(t *testing.T) {
	m := Movement{}
	var err error
	_, err = CreateMovement(m)
	if err == nil {
		t.Fatal("error should have been thrown when name is empty")
	}
}

func TestGetMovement(t *testing.T) {
	var m Movement
	err := GetMovement(ID.Hex(), &m)
	if m.ID != ID {
		t.Fatalf("ID has an incorrect value: %s", m.ID.Hex())
	}
	if m.Name != movementName {
		t.Fatalf("Name has an incorrect value: %s", m.Name)
	}
	if m.Description != movementDescription {
		t.Fatalf("Description has an incorrect value: %s", m.Description)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMovement_NotFound(t *testing.T) {
	var movement Movement
	err := GetMovement("", &movement)
	if err != nil {
		t.Fatal(err)
	}
	if movement != (Movement{}) {
		t.Fatal("expected nil movement")
	}
}

func TestUpdateMovement(t *testing.T) {
	var m Movement
	err := GetMovement(ID.Hex(), &m)
	t.Log(m)
	if err != nil {
		t.Fatal(err)
	}
	m.Name = "AnUpdatedMovementName"
	err = UpdateMovement(m)
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
	if movements[0].Description != movementDescription {
		t.Fatalf("expected movement name to equal: %s but found %s", movementName, movements[0].Name)
	}
}

func TestDeleteMovement(t *testing.T) {
	err := DeleteMovement(ID.Hex())
	if err != nil {
		t.Fatal(err)
	}
}
