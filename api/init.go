package api

import (
	"gopkg.in/mgo.v2"
	"time"
	log "github.com/sirupsen/logrus"
)

const database = "madziki"
const movementsCollection = "movements"

var settings = &mgo.DialInfo{
	Addrs:   []string{"db:27017"},
	Timeout: 30 * time.Second,
	//Database: "madziki",
	//Username: "",
	//Password: "",
}

var session *mgo.Session

func init() {
	log.Debug("Initializing default madziki session...")

	s, err := mgo.DialWithInfo(settings)
	if err != nil {
		log.Panic(err)
	}
	log.Debug("controllers", "init", "Opened a valid db session...")
	session = s

	// Ensure indexes exist.
	log.Debug("Creating madziki indexes...")
	//movementsIndex := mgo.Index{
	//	Key:        []string{"name", "phone"},
	//	Unique:     true,
	//	DropDups:   true,
	//	Background: true,
	//	Sparse:     true,
	//}
	//
	//err = c.EnsureIndex(index)
	//if err != nil {
	//	panic(err)
	//}
}

// GetSession returns the current Mongo session.
func getSession() *mgo.Session {
	return session.Copy()
}
