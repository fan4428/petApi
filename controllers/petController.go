package controllers

import (
	b64 "encoding/base64"
	"fmt"
	"petApi/data"
	"petApi/models"
	"strconv"
	"strings"
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
	var code = c.PostForm("code")
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
	bespeakModels.OpenId = data.GetOpenId("wxf0e257ada269dd09", "9e45db58aea6e7f61e6dc9a53f35f81a", code)
	err := data.InsertBespeak(bespeakModels)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, true)
}

// type KDRespBody struct {
// 	Errcode string `json:"email"`
// }

//Login 登录
func Login(c *gin.Context) {
	// var reqInfo KDRespBody
	// err := c.BindJSON(&reqInfo)
	fan := c.Param("email")
	email := c.PostForm("email")
	password := c.PostForm("password")
	fmt.Println(fan)
	userModel, err := data.EmailLogin(email, password)
	if err != nil {
		fmt.Println(err)
	}
	if userModel.Name != "" {
		var token = string(userModel.ID) + ":" + userModel.Name + ":" + time.Now().Format("2006-01-02 15:04:05")
		text := []byte(token)
		key := []byte("sfe023f_9fd&fwfl")
		dToken, derr := data.Encrypt(text, key)
		if derr != nil {
			panic(derr)
		}
		userModel.Token = b64.StdEncoding.EncodeToString([]byte(dToken))
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

//ValiDateToken 验证token
func ValiDateToken(c *gin.Context) {
	token := c.PostForm("token")
	token = strings.Replace(token, " ", "+", -1)
	decodeBytes, err := b64.StdEncoding.DecodeString(token)
	if err != nil {
		panic(err)
	}
	key := []byte("sfe023f_9fd&fwfl")
	byteToken, derr := data.Dncrypt(decodeBytes, key)
	if derr != nil {
		panic(derr)
	}
	strToken := string(byteToken)
	s := strings.Split(strToken, ":")
	if len(s) >= 3 {
		c.JSON(200, "yes")
	} else {
		c.JSON(200, "no")
	}

}

//GetMyBespeak 获取我的预约
func GetMyBespeak(c *gin.Context) {
	code := c.PostForm("code")
	openid := data.GetOpenId("wxf0e257ada269dd09", "9e45db58aea6e7f61e6dc9a53f35f81a", code)
	//openid = "ohVxV0wvfMjBkN3AMWbOgM6UzEuM"
	bespeakList, derr := data.GetMyBespeak(openid)
	if derr != nil {
		panic(derr)
	}

	c.JSON(200, bespeakList)

}
