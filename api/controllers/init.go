package controllers

import (
	"github.com/blueshirts/madziki/api/logger"
)

var CONTROLLERS = "CONTROLLERS"

func init() {
	logger.Info(CONTROLLERS, "init", "Initializing CONTROLLERS...")
}

//import (
//	mgo "gopkg.in/mgo.v2"
//	"log"
//)

// Setup initializes a redis client
//func Setup() (*mgo.Session, error) {
//	session, err := mgo.Dial("localhost")
//	if err != nil {
//		return nil, err
//	}
//
//	return session, nil
//}
