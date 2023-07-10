package controller

import (
	"Gin_Project/go/entity"
	"Gin_Project/go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserFP(c *gin.Context) {
	temp, _ := c.Params.Get("userid")
	UserID, _ := strconv.Atoi(temp)
	UserFP, err := service.GetFP(UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"data": UserFP,
		})
	}
}

func AddUserFP(c *gin.Context) {
	var temp entity.UserFP
	err := c.BindJSON(&temp)
	if err != nil {
		panic(err)
		return
	}
	var isFind bool
	isFind, err = service.AddFP(temp)
	if err != nil {
		c.JSON(400, gin.H{
			"code":   400,
			"isFind": isFind,
			"data":   entity.UserFP{},
		})
	} else {
		c.JSON(200, gin.H{
			"code":   200,
			"isFind": isFind,
			"data":   temp,
		})
	}
}

func DeleteUserFP(c *gin.Context) {
	var temp entity.UserFP
	err := c.BindJSON(&temp)
	if err != nil {
		panic(err)
		return
	}
	var isFind bool
	isFind, err = service.DeleteFP(temp)
	if err != nil {
		c.JSON(400, gin.H{
			"code":   400,
			"isFind": isFind,
			"data":   entity.UserFP{},
		})
	} else {
		c.JSON(200, gin.H{
			"code":   200,
			"isFind": isFind,
			"data":   temp,
		})
	}
}
