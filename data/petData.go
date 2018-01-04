package data

import (
	b64 "encoding/base64"
	"fmt"
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

//GetAllHospitalByID 通过id获取医院
func GetAllHospitalByID(hid string) (models.Hospital, error) {
	var dbPet = GetDB("Pet")
	cHospital := dbPet.C("hospital")
	ohid := bson.ObjectIdHex(hid)
	defer dbPet.Session.Close()

	var hospitalModel models.Hospital
	errPet := cHospital.Find(bson.M{"_id": ohid}).One(&hospitalModel)

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

	errPet := cBespeak.Insert(&bespeakModel)

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

//FindBespeakFullcalendar 查找一年的所有预约
func FindBespeakFullcalendar(sDate string, eDate string, hospitalID string) ([]models.Bespeak, error) {
	var dbPet = GetDB("Pet")
	cBespeak := dbPet.C("bespeak")
	var bespeakMode []models.Bespeak

	var bsonM = bson.M{"hospitalId": hospitalID, "applyDate": bson.M{"$gte": sDate, "$lte": eDate}}

	errPet := cBespeak.Find(bsonM).All(&bespeakMode)
	return bespeakMode, errPet

}

//EmailLogin 登录系统
func EmailLogin(email string, password string) (models.User, error) {
	var dbPet = GetDB("ApplicationCentre")
	cUser := dbPet.C("user")

	text := []byte(password)
	key := []byte("sfe023f_9fd&fwfl")
	dPassword, derr := Encrypt(text, key)
	uEnc := b64.StdEncoding.EncodeToString([]byte(dPassword))
	fmt.Println("uEncerr:", uEnc)

	if derr != nil {
		panic(derr)
	}
	var userModel models.User
	var bsonM = bson.M{"loginId": email, "password": uEnc}

	errUser := cUser.Find(&bsonM).One(&userModel)
	return userModel, errUser
}

//GetMyBespeak 获取我的预约
func GetMyBespeak(openid string) ([]models.FindBespeak, error) {
	var dbPet = GetDB("Pet")
	cBespeak := dbPet.C("bespeak")

	var bespeakMode []models.FindBespeak

	var bsonM = bson.M{"openid": openid}

	errPet := cBespeak.Find(bsonM).Sort("-applyDate", "applyTime").Limit(10).All(&bespeakMode)
	for index := 0; index < len(bespeakMode); index++ {
		hid := bespeakMode[index].HospitalID
		hospital, herr := GetAllHospitalByID(hid)
		if herr != nil {
			panic(herr)
		}
		bespeakMode[index].Hospital = hospital

	}

	return bespeakMode, errPet
}
