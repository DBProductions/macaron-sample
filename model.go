package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Session

const mongodbUri = "mongodb://mongodb/golang"

type Person struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Age  int
	Time string
}

var (
	mgoSession *mgo.Session
)

func ConnectDb() {
	if mgoSession == nil {
		session, err := mgo.Dial(mongodbUri)
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
		mgoSession = session
	}
	return
}

func GetDbSession() *mgo.Session {
	if mgoSession == nil {
		ConnectDb()
	}
	return mgoSession
}

func CloseDb() {
	mgoSession.Close()
	return
}
