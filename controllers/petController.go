package controllers

import (
	"fmt"
	"petApi/data"
	"petApi/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

//GetAllHospital 获取所有医院
func GetAllHospital(c *gin.Context) {
	result, err := data.GetAllHospital()
	if err != nil {
		fmt.Println(err)
	}
	// for i := 0; i < result.length; i++ {
	// 	for j := 0; j < result[i].Department; j++ {
	// 		result[i].Department[j].Doctor
	// 	}

	// }
	// doctorResult,err:=data.GetDoctorByID()
	c.JSON(200, result)
}

//GetDoctorByID 通过doctorId获取医生
func GetDoctorByID(c *gin.Context) {
	var doctorID = c.Query("doctorId")
	result, err := data.GetDoctorByID(doctorID)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}

//InsertBespeak 添加预约
func InsertBespeak(c *gin.Context) {
	var bespeakModels models.Bespeak
	t := time.Now()
	bespeakModels.ID = bson.NewObjectId()
	bespeakModels.MasterName = c.PostForm("masterName")
	bespeakModels.PetName = c.PostForm("petName")
	bespeakModels.Mobile = c.PostForm("mobile")
	bespeakModels.ApplyDate = c.PostForm("applyDate")
	bespeakModels.ApplyTime = c.PostForm("applyTime")
	bespeakModels.HospitalID = c.PostForm("hospitalId")
	bespeakModels.HospitalName = c.PostForm("hospitalName")
	bespeakModels.DepID = c.PostForm("depId")
	bespeakModels.DepName = c.PostForm("depName")
	bespeakModels.DoctorID = c.PostForm("doctorId")
	bespeakModels.DoctorName = c.PostForm("doctorName")
	bespeakModels.State = 0
	bespeakModels.CreateDate = t
	err := data.InsertBespeak(bespeakModels)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, true)
}

//Login 登录
func Login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	userModel, err := data.EmailLogin(email, password)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, userModel)
}

//FindBespeak 查询预约
func FindBespeak(c *gin.Context) {

	// var bespeakModels models.Bespeak

	applyDate := c.PostForm("applyDate")
	pageIndex := c.PostForm("pageIndex")
	pageCount := c.PostForm("pageCount")

	iPageIndex, errIndex := strconv.Atoi(pageIndex)
	iPageCount, errCount := strconv.Atoi(pageCount)

	if errIndex != nil && errCount != nil {

	}
	bespeakModels, err := data.FindBespeak(applyDate, iPageIndex, iPageCount)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, bespeakModels)
}

//FindBespeakFullcalenar 查询预约
func FindBespeakFullcalenar(c *gin.Context) {

	// var bespeakModels models.Bespeak

	sDate := c.PostForm("sDate")
	eDate := c.PostForm("eDate")
	hid := c.PostForm("hid")

	bespeakModels, err := data.FindBespeakFullcalendar(sDate, eDate, hid)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, bespeakModels)
}
