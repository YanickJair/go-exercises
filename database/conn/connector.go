package conn

import (
	"log"
	"sync"

	"gopkg.in/mgo.v2"
)

var database *mgo.Database
var once sync.Once

// Connector - connect to our db and make sure we connect only once
func Connector(url, dbname string) *mgo.Database {
	once.Do(func() {
		session := makeSession(url)
		database = session.DB(dbname)
	})
	return database
}

// makeSession - connect to a URL and return
func makeSession(url string) *mgo.Session {
	session, err := mgo.Dial(url)
	//! Check if any error occurred
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	//! Make sure we allways close our db on error events
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	return session
}
