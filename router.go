package main

import (
	"log"
	"petApi/controllers"

	"petApi/data"

	"github.com/gin-gonic/gin"
)

//initRouter 初始化路由
func initRouter() *gin.Engine {

	data.GetOpenId("wxf0e257ada269dd09", "9e45db58aea6e7f61e6dc9a53f35f81a", "071REjr12W4VjX02uep12lVgr12REjrI")
	//go petglobal.GetAccessTokenTimer("wxf0e257ada269dd09", "9e45db58aea6e7f61e6dc9a53f35f81a", 10)
	//petglobal.GetAccessToken("wxf0e257ada269dd09", "9e45db58aea6e7f61e6dc9a53f35f81a")
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("content-type", "application/json")

		log.Println("all")
		c.Next()
	})

	//Pet接口
	pet := router.Group("Petapi/v1")
	{
		pet.OPTIONS("/insertbespeak", func(c *gin.Context) {
			c.JSON(200, "")
		})
		pet.GET("/hospital", controllers.GetAllHospital)
		pet.GET("/getdoctorbyid", controllers.GetDoctorByID)
		pet.POST("/insertbespeak", controllers.InsertBespeak)
		pet.POST("/findBespeak", controllers.FindBespeak)
		pet.POST("/emailLogin", controllers.Login)
		pet.POST("/findBespeakFullcalenar", controllers.FindBespeakFullcalenar)
		pet.POST("/ValiDateToken", controllers.ValiDateToken)
		pet.POST("/GetMyBespeak", controllers.GetMyBespeak)
	}
	//ceshi
	v1 := router.Group("api/v1")
	{
		v1.GET("/users", controllers.GetUserByname)
		//	v1.GET("/test", controllers.Test)

	}
	//测试接口
	test := router.Group("api/test")
	{

		test.POST("/test", controllers.InsertTest)

	}

	return router
}
