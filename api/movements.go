package api

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	log "github.com/sirupsen/logrus"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

var validate = validator.New()

// Movement is a encapsulation of a single movement instance.
type Movement struct {
	ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `json:"name" validate:"required"`
	Description string        `json:"description"`
	Details     string        `json:"details"`
	User        string        `json:"user" validate:"required"`
	Created     time.Time     `json:"created" validate:"required"`
	Updated     time.Time     `json:"updated" validate:"required"`
}

func ValidateMovement(m Movement) error {
	return validate.Struct(m)
}

// Create movement creates a new movement object in the database.
func CreateMovement(movement Movement) (bson.ObjectId, error) {
	err := ValidateMovement(movement)
	if err != nil {
		return "", err
	}
	session := getSession()
	defer session.Close()

	session.SetSafe(&mgo.Safe{})

	db := session.DB(database)
	movements := db.C(movementsCollection)

	movement.ID = bson.NewObjectId()
	err = movements.Insert(&movement)
	if err != nil {
		return "", err
	}

	log.Debugf("Created new movement with id: %s", movement.ID)
	return movement.ID, err
}

func GetMovement(hex string, movement *Movement) error {
	var objectId bson.ObjectId
	err := getObjectId(hex, &objectId)
	if err != nil {
		// exit, could not find a related movement
		return nil
	}

	session := getSession()
	defer session.Close()

	db := session.DB(database)
	movements := db.C(movementsCollection)
	return movements.Find(bson.M{"_id": objectId}).One(movement)
}

func UpdateMovement(m Movement) error {
	err := ValidateMovement(m)
	if err != nil {
		return err
	}
	session := getSession()
	defer session.Close()

	session.SetSafe(&mgo.Safe{})

	db := session.DB(database)
	movements := db.C(movementsCollection)
	filter := bson.M{"_id": m.ID}
	return movements.Update(filter, &m)
}

// DeleteMovement deletes and existing movement by movement.
func DeleteMovement(hex string) error {
	var objectId bson.ObjectId
	err := getObjectId(hex, &objectId)
	if err != nil {
		return errors.New(fmt.Sprintf("invalid hex format: %s", hex))
	}

	session := getSession()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	db := session.DB(database)
	return db.C(movementsCollection).Remove(bson.M{"_id": objectId})
}

// ListMovements retrieves all matching movements from the collection.
func ListMovements(movements *[]Movement) error {
	session := getSession()
	defer session.Close()
	db := session.DB(database)
	return db.C(movementsCollection).Find(nil).All(movements)
}
