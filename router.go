package main

import (
	"log"
	"petApi/controllers"

	"github.com/gin-gonic/gin"
)

//initRouter 初始化路由
func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("content-type", "application/json")

		log.Println("all")
	})

	//Pet接口
	pet := router.Group("api/pet")
	{
		pet.GET("/hospital", controllers.GetAllHospital)

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
