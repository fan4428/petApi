package data

import (
	"github.com/widuu/goini"
	mgo "gopkg.in/mgo.v2"
)

//GetDB 获取数据库
func GetDB(dbName string) *mgo.Database {

	conf := goini.SetConfig("./config/conf.ini")
	dbPath := conf.GetValue("mongo", "dbPath")

	session, err := mgo.Dial(dbPath)
	if err != nil {
		panic(err)
	}
	db := session.DB(dbName)

	return db
}
