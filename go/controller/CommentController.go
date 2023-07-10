package controller

import (
	"Gin_Project/go/entity"
	"Gin_Project/go/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Comment struct {
	PostID    int `json:"PostID"`
	CommentID int `json:"CommentID"`
}

func PostPage(c *gin.Context) {
	session := sessions.Default(c)
	User, _ := session.Get("User").(entity.User)
	c.JSON(200, gin.H{
		"data": []*entity.Comment{},
		"user": User,
	})
}

func CreateComment(c *gin.Context) {
	//定义一个User变量
	var comment entity.Comment
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	err := c.BindJSON(&comment)
	comment.Time = time.Now().Format("2006-01-02 15:04:05")
	comment.UpNumber = 0
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err = service.CreateComment(&comment)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": comment,
		})
	}
}

func GetPostComment(c *gin.Context) {
	session := sessions.Default(c)
	User, _ := session.Get("User").(entity.User)
	PostID, _ := c.Params.Get("id")
	id, _ := strconv.Atoi(PostID)
	todoList, err := service.GetPostComment(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"data":   todoList,
			"PostID": id,
			"user":   User,
		})
	}
}

func TimeSort(c *gin.Context) {
	session := sessions.Default(c)
	User, _ := session.Get("User").(entity.User)
	PostID, _ := c.Params.Get("id")
	id, _ := strconv.Atoi(PostID)
	todoList, err := service.TimeSort(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"data":   todoList,
			"PostID": id,
			"user":   User,
		})
	}
}

func Like(c *gin.Context) {
	var temp Comment
	err := c.BindJSON(&temp)
	if err != nil {
		return
	}
	var comment entity.Comment
	if err := service.LikeComment(temp.PostID, temp.CommentID, &comment); err != nil {
		err.Error()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Update success",
			"data": comment,
		})
	}
}

func DisLike(c *gin.Context) {
	var temp Comment
	err := c.BindJSON(&temp)
	if err != nil {
		return
	}
	var comment entity.Comment
	if err := service.DisLikeComment(temp.PostID, temp.CommentID, &comment); err != nil {
		err.Error()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Update success",
			"data": comment,
		})
	}
}

func DeleteCommentById(c *gin.Context) {
	PostID, _ := c.Params.Get("post_id")
	CommentID, _ := c.Params.Get("comment_id")
	p, _ := strconv.Atoi(PostID)
	com, _ := strconv.Atoi(CommentID)
	if err := service.DeleteCommentById(p, com); err != nil {
		err.Error()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": p,
		})
	}
}
