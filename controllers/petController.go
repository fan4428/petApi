package controllers

import "github.com/gin-gonic/gin"
import "petApi/data"
import "fmt"

//GetAllHospital 获取所有医院
func GetAllHospital(c *gin.Context) {
	result, err := data.GetAllHospital()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, result)
}
