package controller

import (
	"Gin_Project/go/common/rsp"
	"Gin_Project/go/entity"
	"Gin_Project/go/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateSpot1(c *gin.Context) {
	token := c.PostForm("token")
	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	sid := c.PostForm("sid")
	n1, _ := strconv.Atoi(sid)

	//定义一个Spot变量
	var spot entity.Spot1
	spot.Uid = n
	spot.Sid = n1

	if err := service.CreateSpot1(&spot); err != nil {
		rsp.Error(c, "收藏失败")
	} else {
		rsp.Success(c, "收藏成功", spot)
	}
}

func GetSpotList1(c *gin.Context) {
	if todoList, err := service.GetAllSpot1(); err != nil {
		rsp.Error(c, "显示景点收藏列表失败")
	} else {
		rsp.Success(c, "显示景点收藏列表成功", todoList)
	}
}

func DeleteSpot1(c *gin.Context) {
	token := c.PostForm("token")
	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	sid := c.PostForm("sid")
	n1, _ := strconv.Atoi(sid)

	if err := service.DeleteSpotById1(n, n1); err != nil {
		rsp.Error(c, "取消收藏失败")
	} else {
		rsp.Success(c, "取消收藏成功", sid)
	}
}

func GetSpot1ById(uid int, sid int) (a int) {
	//查询景区收藏
	_, err := service.GetSpotById1(uid, sid)
	if err != nil {
		return 0
	}
	return 1
}
