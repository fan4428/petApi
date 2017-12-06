package models

import (
	"gopkg.in/mgo.v2/bson"
)

//User fan
type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Password  string        `bson:"password"`
	TermFlag  bool          `bson:"termflag"`
	AdminRole string        `bson:"AdminRole"`
}

//Result fanhui
type Result struct {
	Code    int
	Message string
	Data    bool
}
