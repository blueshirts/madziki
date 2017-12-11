package api

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	log "github.com/sirupsen/logrus"
)

// Movement is a encapsulation of a single movement instance.
type Movement struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
	Details     string
	User        string
}

// Create movement creates a new movement object in the database.
func CreateMovement(movement Movement) (bson.ObjectId, error) {
	log.Debugf("Movement name: %s", movement.Name)
	if movement.Name == "" {
		return "", errors.New("no movement name is defined")
	}
	session := getSession()
	defer session.Close()

	session.SetSafe(&mgo.Safe{})

	db := session.DB(database)
	movements := db.C(movementsCollection)

	movement.ID = bson.NewObjectId()
	err := movements.Insert(&movement)
	if err != nil {
		log.Info("created new movement with id: %s", movement.ID.Hex())
	}
	return movement.ID, err
}

func GetMovement(id bson.ObjectId, movement *Movement) error {
	session := getSession()
	defer session.Close()

	db := session.DB(database)
	movements := db.C(movementsCollection)
	return movements.Find(bson.M{"_id": id}).One(movement)
}

// DeleteMovement deletes and existing movement by movement.
func DeleteMovement(id bson.ObjectId) error {
	if id == "" {
		errors.New("no movement id is defined")
	}
	session := getSession()
	defer session.Close()

	session.SetSafe(&mgo.Safe{})

	db := session.DB(database)
	return db.C(movementsCollection).Remove(bson.M{"_id": id})
}

// ListMovements retrieves all matching movements from the collection.
func ListMovements(movements *[]Movement) error {
	session := getSession()
	defer session.Close()
	db := session.DB(database)
	return db.C(movementsCollection).Find(nil).All(movements)
}
