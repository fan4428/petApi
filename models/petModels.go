package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Hospital 医院包含职位
type Hospital struct {
	ID             bson.ObjectId `bson:"_id"`
	HospitalName   string        `bson:"hospitalName"`
	HospitalRemark string        `bson:"hospitalRemark"`
	Department     []Department  `bson:"department"`
	State          int           `bson:"state"`
	Address        string        `bson:"address"`
	Doctor         []Doctor      `bson:"doctor"`
	Location       Location      `bson:"location"`
}

//坐标位置
type Location struct {
	Lng float64 `bson:"lng"`
	Lat float64 `bson:"lat"`
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
	Remark     string        `bson:"Remark"`
}

//Bespeak 预约表
type Bespeak struct {
	ID           bson.ObjectId `bson:"_id"`
	MasterName   string        `bson:"masterName"`
	PetName      string        `bson:"petName"`
	Mobile       string        `bson:"mobile"`
	ApplyDate    string        `bson:"applyDate"`
	ApplyTime    string        `bson:"applyTime"`
	HospitalID   string        `bson:"hospitalId"`
	HospitalName string        `bson:"hospitalName"`
	DepID        string        `bson:"depId"`
	DepName      string        `bson:"depName"`
	DoctorID     string        `bson:"doctorId"`
	DoctorName   string        `bson:"doctorName"`
	State        int           `bson:"state"`
	CreateDate   time.Time     `bson:"createDate"`
	OpenId       string        `bson:"openid"`
}

//FindBespeak 预约表
type FindBespeak struct {
	ID           bson.ObjectId `bson:"_id"`
	MasterName   string        `bson:"masterName"`
	PetName      string        `bson:"petName"`
	Mobile       string        `bson:"mobile"`
	ApplyDate    string        `bson:"applyDate"`
	ApplyTime    string        `bson:"applyTime"`
	HospitalID   string        `bson:"hospitalId"`
	HospitalName string        `bson:"hospitalName"`
	DepID        string        `bson:"depId"`
	DepName      string        `bson:"depName"`
	DoctorID     string        `bson:"doctorId"`
	DoctorName   string        `bson:"doctorName"`
	State        int           `bson:"state"`
	CreateDate   time.Time     `bson:"createDate"`
	OpenId       string        `bson:"openid"`
	Hospital     Hospital
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
