package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/macaron.v1"
	"labix.org/v2/mgo/bson"
	"time"
)

func overviewHandler(ctx *macaron.Context) {
	db := GetDbSession()
	c := db.DB("golang").C("people")
	var results []Person
	err := c.Find(bson.M{}).All(&results)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, &results)
}

func detailHandler(ctx *macaron.Context) {
	db := GetDbSession()
	c := db.DB("golang").C("people")
	result := Person{}
	err := c.FindId(bson.ObjectIdHex(ctx.Params("id"))).One(&result)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"uri":    ctx.Req.RequestURI,
			"method": ctx.Req.Method,
		}).Error("Error")
		ctx.JSON(404, map[string]string{"message": "not found", "id": ctx.Params("id")})
	} else {
		ctx.JSON(200, &result)
	}
}

func createHandler(ctx *macaron.Context, person Person) {
	db := GetDbSession()
	c := db.DB("golang").C("people")
	t := time.Now()
	id := bson.NewObjectId()
	err := c.Insert(&Person{ID: id, Name: person.Name, Age: person.Age, Time: fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())})
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"uri":    ctx.Req.RequestURI,
			"method": ctx.Req.Method,
		}).Error("Error")
		ctx.JSON(400, map[string]string{"message": "error"})
	} else {
		ctx.JSON(201, map[string]string{"message": "Created", "id": id.String()})
	}
}

func updateHandler(ctx *macaron.Context, person Person) {
	db := GetDbSession()
	c := db.DB("golang").C("people")
	change := bson.M{"$set": bson.M{"name": person.Name, "age": person.Age, "time": time.Now()}}
	err := c.UpdateId(bson.ObjectIdHex(ctx.Params("id")), change)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"uri":    ctx.Req.RequestURI,
			"method": ctx.Req.Method,
		}).Error("Error")
		ctx.JSON(404, map[string]string{"message": "not found", "id": ctx.Params("id")})
	}
	ctx.JSON(201, map[string]string{"message": "updated", "id": ctx.Params("id")})
}

func upgradeHandler(ctx *macaron.Context, person Person) {
	db := GetDbSession()
	c := db.DB("golang").C("people")
	change := bson.M{"$set": bson.M{"name": person.Name, "age": person.Age, "time": time.Now()}}
	err := c.UpdateId(bson.ObjectIdHex(ctx.Params("id")), change)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"uri":    ctx.Req.RequestURI,
			"method": ctx.Req.Method,
		}).Error("Error")
		ctx.JSON(404, map[string]string{"message": "not found", "id": ctx.Params("id")})
	}
	ctx.JSON(201, map[string]string{"message": "upgraded", "id": ctx.Params("id")})
}

func deleteHandler(ctx *macaron.Context) {
	db := GetDbSession()
	c := db.DB("golang").C("people")
	err := c.RemoveId(bson.ObjectIdHex(ctx.Params("id")))
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"uri":    ctx.Req.RequestURI,
			"method": ctx.Req.Method,
		}).Error("Error")
		ctx.JSON(404, map[string]string{"message": "not found", "id": ctx.Params("id")})
	} else {
		ctx.JSON(204, map[string]string{"message": "deleted", "id": ctx.Params("id")})
	}
}
