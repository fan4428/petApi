package controllers

import (
	"fmt"
	"petApi/data"

	"github.com/gin-gonic/gin"

	model "petApi/models"
)

//GetUserByname 获取user
func GetUserByname(c *gin.Context) {
	var result = make(chan []model.User)

	go data.GetUserByName("Colin CHIU", result)
	v := <-result
	fmt.Println(v)
}
