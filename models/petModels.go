package models

import "gopkg.in/mgo.v2/bson"

//Hospital 医院包含职位
type Hospital struct {
	ID             bson.ObjectId `bson:"_id"`
	HospitalName   string        `bson:"hospitalName"`
	HospitalRemark string        `bson:"hospitalRemark"`
	Department     []Department  `bson:"department"`
	State          int           `bson:"state"`
	Address        string        `bson:"address"`
}

//Department 职位包含doctor数组oid
type Department struct {
	ID        bson.ObjectId   `bson:"depId"`
	DepName   string          `bson:"depName"`
	DepRemark string          `bson:"depRemark"`
	Doctor    []bson.ObjectId `bson:"doctor"`
	state     int             `bson:"state"`
}

//Doctor 医生
type Doctor struct {
	ID         bson.ObjectId `bson:"_id"`
	DoctorName string        `bson:"doctorName"`
}

// "_id" : ObjectId("5a1cc2be080e9befff6bcbec"),
// "hospitalName" : "总院",
// "hospitalRemark" : "全都擅长",
// "department" : [
// 	{
// 		"depId" : ObjectId("521cc2be080e9befff6bcbec"),
// 		"depName" : "内科",
// 		"depRemark" : "内科治疗",
// 		"doctor" : [
// 			ObjectId("5a1cc2fb080e9befff6bcbed"),
// 			ObjectId("5a1cc98f080e9befff6bcbef")
// 		],
// 		"state" : 0
// 	}
// ],
// "address" : "农大abc路123号",
// "state" : 0