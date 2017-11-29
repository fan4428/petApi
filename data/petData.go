package data

import (
	"petApi/models"

	"gopkg.in/mgo.v2/bson"
)

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

//GetAllHospital 获取全部所有医院
func GetAllHospital() ([]models.Hospital, error) {
	var db = GetDB("Pet")
	c := db.C("hospital")

	defer db.Session.Close()

	var hospitalModel []models.Hospital
	err := c.Find(bson.M{}).All(&hospitalModel)

	return hospitalModel, err
}
