package data

import (
	"github.com/widuu/goini"
	mgo "gopkg.in/mgo.v2"
)

//GetDB 获取数据库
func GetDB(dbName string) *mgo.Database {

	conf1 := goini.SetConfig("./config/conf.ini")
	dbPath1 := conf1.GetValue("mongo", "dbPath")

	session, errDB := mgo.Dial(dbPath1)
	session.SetMode(mgo.Monotonic, true)
	if errDB != nil {
		panic(errDB)
	}
	db := session.DB(dbName)

	return db
}
