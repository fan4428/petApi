package controllers

import (
	data "vconnectservice/data"

	"github.com/gin-gonic/gin"
)

//Insert fan
func Insert(c *gin.Context) {
	go data.Insrert()
}
