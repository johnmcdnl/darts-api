package data

import (
	"gopkg.in/mgo.v2"
	"log"
	"sync"
)

const dbName string = "darts"

var once sync.Once
var db *mgo.Database

func mongoSession() *mgo.Session {
	session, err := mgo.Dial("mongo")
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func resourcesDatabase() *mgo.Database {
	return mongoSession().DB(dbName)
}

func Database() *mgo.Database {
	once.Do(func() {
		db = resourcesDatabase()
	})
	return db
}