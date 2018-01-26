package api

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type system struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Details     string        `json:"details"`
	Created     time.Time     `json:"created"`
	Updated     time.Time     `json:"updated"`
}
