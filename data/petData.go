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
	var dbPet = GetDB("Pet")
	cHospital := dbPet.C("hospital")

	defer dbPet.Session.Close()

	var hospitalModel []models.Hospital
	errPet := cHospital.Find(bson.M{}).All(&hospitalModel)

	return hospitalModel, errPet
}

//GetDoctorByID 通过doctorId获取医生
func GetDoctorByID(doctorID string) (models.Doctor, error) {

	var bsonDoctorID = bson.ObjectIdHex(doctorID)

	var dbPet = GetDB("Pet")
	cHospital := dbPet.C("doctor")
	defer dbPet.Session.Close()
	var doctorModel models.Doctor
	errPet := cHospital.Find(bson.M{"_id": bsonDoctorID}).One(&doctorModel)

	return doctorModel, errPet
}

//InsertBespeak 添加预约
func InsertBespeak(bespeakModel models.Bespeak) error {
	var dbPet = GetDB("Pet")
	cBespeak := dbPet.C("bespeak")
	defer dbPet.Session.Close()

	errPet := cBespeak.Insert(bespeakModel)

	return errPet
}

//FindBespeak 查找预约
func FindBespeak(applyDate string, skip int, limit int) ([]models.Bespeak, error) {
	var dbPet = GetDB("Pet")
	cBespeak := dbPet.C("bespeak")
	var bespeakMode []models.Bespeak

	var bsonM = bson.M{"applyDate": applyDate}
	pskip := (skip - 1) * limit
	errPet := cBespeak.Find(bsonM).Skip(pskip).Limit(limit).All(&bespeakMode)
	return bespeakMode, errPet
}
