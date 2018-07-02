package database

import (
	"gopkg.in/mgo.v2"
)

var database *mgo.Session

func Initialize() error {
	db, err := mgo.Dial("localhost")
	
	if err != nil {
		return err
	} 

	database = db
	return nil
}

func Instance() *mgo.Session {
	return database.Clone()
}