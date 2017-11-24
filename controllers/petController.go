package controllers

import (
	data "petApi/data"
	model "petApi/models"

	"github.com/gin-gonic/gin"
)

//InsertTest 获取user
func InsertTest(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("content-type", "application/json")

	name := c.PostForm("name")
	gender := c.PostForm("gender")
	email := c.PostForm("email")
	hukou := c.PostForm("hukou")
	englishname := c.PostForm("englishname")
	birthyear := c.PostForm("birthyear")
	mobile := c.PostForm("mobile")
	living := c.PostForm("living")
	educationlevel := c.PostForm("educationlevel")
	university := c.PostForm("university")
	major := c.PostForm("major")
	currentsalary := c.PostForm("currentsalary")
	reason := c.PostForm("reason")
	expectedsalary := c.PostForm("expectedsalary")

	data.InsertTest(name, gender, email, hukou,
		englishname, birthyear, mobile, living, educationlevel, university, major,
		currentsalary, reason, expectedsalary)
	result := &model.Result{}
	result.Code = 0
	result.Message = ""
	result.Data = true

	c.JSON(200, result)

}
