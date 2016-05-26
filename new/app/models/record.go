package models

import (
	"labix.org/v2/mgo/bson"
)

type User struct {
	id       bson.ObjectId `bson:"_id"`
	name     string        `bson:"Name"`
	password string        `bson:"Passwd"`
}
type Doc struct {
	id      bson.ObjectId `bson:"_id"`
	name    string        `bson:"Name"`
	tag     string        `bson:"Tag"`
	contain string        `bson:Contain`
	Com     []string      `bson:"com"`
}
