package main

import (
	"vconnectservice/controllers"

	"github.com/gin-gonic/gin"
)

//initRouter 初始化
func initRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/users", controllers.Insert)

	}

	return router
}
