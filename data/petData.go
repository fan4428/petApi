package data

import (
	"gopkg.in/mgo.v2/bson"
)

//F fan

// //AdminLogin 宠物后台系统登录
// func AdminLogin(user string, pwd string) {
// 	c := db.C("user")

// 	err := c.Find(b).One(bson.M{"name": "fan"})
// 	defer db.Session.Close()

// }

// InsertTest 通过名称获取User用户
func InsertTest(name string, gender string, email string, hukou string,
	englishname string, birthyear string, mobile string, living string, educationlevel string,
	university string, major string, currentsalary string, reason string, expectedsalary string) {
	var db = GetDB("Pet")
	defer db.Session.Close()
	c := db.C("test")

	err := c.Insert(bson.M{"name": name, "gender": gender,
		"email": email, "hukou": hukou,
		"englishname": englishname, "birthyear": birthyear,
		"mobile": mobile, "living": living,
		"educationlevel": educationlevel, "university": university,
		"major": major, "currentsalary": currentsalary,
		"reason": reason, "expectedsalary": expectedsalary})
	if err != nil {
		panic(err)
	}

}
