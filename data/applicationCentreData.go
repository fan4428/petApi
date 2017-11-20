package data

import (
	"petApi/models"

	"gopkg.in/mgo.v2/bson"
)

//F fan

// //AdminLogin 宠物后台系统登录
// func AdminLogin(user string, pwd string) {
// 	c := db.C("user")

// 	err := c.Find(b).One(bson.M{"name": "fan"})
// 	defer db.Session.Close()

// }

// GetUserByName 通过名称获取User用户
func GetUserByName(userName string, result chan []models.User) {
	var db = GetDB("ApplicationCentre")
	defer db.Session.Close()
	c := db.C("user")

	var uModel []models.User
	err := c.Find(bson.M{"name": userName}).All(&uModel)
	if err != nil {
		panic(err)
	}
	result <- uModel

}
