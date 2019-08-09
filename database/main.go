package main

import (
	"fmt"
	"log"

	"github.com/YanickJair/go-exercises/database/conn"
	"gopkg.in/mgo.v2/bson"
)

// Person - a human person object
type Person struct {
	id    string `bson:"_id"`
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
}

func main() {
	db := conn.Connector("mongodb://localhost", "dummy-auth")

	c := db.C("people")
	err := c.Insert(&Person{Name: "Djamila", Phone: "999999"},
		&Person{Name: "Gabriel", Phone: "5975199"})
	if err != nil {
		log.Fatal("Could not insert in DB")
	}
	res := Person{}
	err = c.Find(bson.M{"name": "Gabriel"}).One(&res)
	if err != nil {
		log.Fatal("Person not found")
	}
	fmt.Println("Person found was: ", res)
}
