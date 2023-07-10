package controller

import (
	"Gin_Project/go/entity"
	"Gin_Project/go/service"
	"github.com/gin-gonic/gin"
	_ "github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	PostID int `json:"PostID"`
}

func PostsPage(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func CreatePost(c *gin.Context) {
	var post entity.Post
	post.UserID, _ = strconv.Atoi(c.PostForm("userID"))
	post.Title = c.PostForm("title")
	post.Tag1 = c.PostForm("tag1")
	post.Tag2 = c.PostForm("tag2")
	post.Contents = c.PostForm("content")

	if c.Request.MultipartForm.File["fileList[0][raw]"] != nil {
		fileHeader := c.Request.MultipartForm.File["fileList[0][raw]"]
		file, err := fileHeader[0].Open()
		defer file.Close()
		if err != nil {
			return
		}
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			return
		}
		post.Photo = fileData
	}

	if c.Request.MultipartForm.File["fileList[1][raw]"] != nil {
		fileHeader1 := c.Request.MultipartForm.File["fileList[1][raw]"]
		file1, err := fileHeader1[0].Open()
		defer file1.Close()
		if err != nil {
			return
		}
		fileData1, err := ioutil.ReadAll(file1)
		if err != nil {
			return
		}
		post.Photo1 = fileData1
	}

	if c.Request.MultipartForm.File["fileList[2][raw]"] != nil {
		fileHeader2 := c.Request.MultipartForm.File["fileList[2][raw]"]
		file2, err := fileHeader2[0].Open()
		defer file2.Close()
		if err != nil {
			return
		}
		fileData2, err := ioutil.ReadAll(file2)
		if err != nil {
			return
		}
		post.Photo2 = fileData2
	}

	if c.Request.MultipartForm.File["fileList[3][raw]"] != nil {
		fileHeader3 := c.Request.MultipartForm.File["fileList[3][raw]"]
		file3, err := fileHeader3[0].Open()
		defer file3.Close()
		if err != nil {
			return
		}
		fileData3, err := ioutil.ReadAll(file3)
		if err != nil {
			return
		}
		post.Photo3 = fileData3
	}

	if c.Request.MultipartForm.File["fileList[4][raw]"] != nil {
		fileHeader4 := c.Request.MultipartForm.File["fileList[4][raw]"]
		file4, err := fileHeader4[0].Open()
		defer file4.Close()
		if err != nil {
			return
		}
		fileData4, err := ioutil.ReadAll(file4)
		if err != nil {
			return
		}
		post.Photo4 = fileData4
	}

	if c.Request.MultipartForm.File["fileList[5][raw]"] != nil {
		fileHeader5 := c.Request.MultipartForm.File["fileList[5][raw]"]
		file5, err := fileHeader5[0].Open()
		defer file5.Close()
		if err != nil {
			return
		}
		fileData5, err := ioutil.ReadAll(file5)
		if err != nil {
			return
		}
		post.Photo5 = fileData5
	}

	if c.Request.MultipartForm.File["fileList[6][raw]"] != nil {
		fileHeader6 := c.Request.MultipartForm.File["fileList[6][raw]"]
		file6, err := fileHeader6[0].Open()
		defer file6.Close()
		if err != nil {
			return
		}
		fileData6, err := ioutil.ReadAll(file6)
		if err != nil {
			return
		}
		post.Photo6 = fileData6
	}

	if c.Request.MultipartForm.File["fileList[7][raw]"] != nil {
		fileHeader7 := c.Request.MultipartForm.File["fileList[7][raw]"]
		file7, err := fileHeader7[0].Open()
		defer file7.Close()
		if err != nil {
			return
		}
		fileData7, err := ioutil.ReadAll(file7)
		if err != nil {
			return
		}
		post.Photo7 = fileData7
	}

	if c.Request.MultipartForm.File["fileList[8][raw]"] != nil {
		fileHeader8 := c.Request.MultipartForm.File["fileList[8][raw]"]
		file8, err := fileHeader8[0].Open()
		defer file8.Close()
		if err != nil {
			return
		}
		fileData8, err := ioutil.ReadAll(file8)
		if err != nil {
			return
		}
		post.Photo8 = fileData8
	}

	post.RTime = time.Now().Format("2006/01/02 15:04")

	err := service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": post,
		})
	}
}
func GetPostList(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	postlist, err := service.GetAllPost()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": postlist,
		})
	}
}
func GetPosts(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	postlist, err := service.GetPosts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": postlist,
		})
	}
}
func GetUserID(c *gin.Context) {
	c.HTML(http.StatusOK, "SearchByUserID.html", gin.H{
		"code": 200,
		"msg":  "success",
	})
}
func GetMyPostList(c *gin.Context) {
	userID, _ := c.Params.Get("id")
	id, _ := strconv.Atoi(userID)
	mypost, err := service.GetMyPost(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title": "My Posts",
			"posts": mypost,
		})
	}
}
func EditMyPost(c *gin.Context) {
	postID := c.Param("postid")
	id, _ := strconv.Atoi(postID)
	editpost, err := service.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title": "Edit My Post",
			"post":  editpost,
		})
	}

}
func UpdatePost(c *gin.Context) {
	postID, _ := c.Params.Get("postid")
	id, _ := strconv.Atoi(postID)
	newEditpost, err := service.GetPostByID(id)
	var temp entity.Post
	err = c.BindJSON(&temp)
	if err != nil {
		print(err)
	}
	newEditpost[0].Title = temp.Title
	newEditpost[0].Contents = temp.Contents
	newEditpost[0].Photo = temp.Photo
	newEditpost[0].Tag1 = temp.Tag1
	newEditpost[0].Tag2 = temp.Tag2
	newEditpost[0].RTime = time.Now().Format("2006/01/02 15:04")
	err = service.UpdatePost(id, newEditpost[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": newEditpost[0],
		})
	}
}
func DeletePost(c *gin.Context) {
	post, _ := c.Params.Get("postid")
	data, _ := strconv.Atoi(post)
	if err := service.DeletePostByID(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": post,
		})
	}
}
func LikePost(c *gin.Context) {
	var postid Post
	err := c.BindJSON(&postid)
	if err != nil {
		return
	}
	var post entity.Post
	if err := service.LikePost(postid.PostID, &post); err != nil {
		err.Error()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Update success",
			"data": post,
		})
	}
}
func DislikePost(c *gin.Context) {
	var postid Post
	err := c.BindJSON(&postid)
	if err != nil {
		return
	}
	var post entity.Post
	if err := service.DislikePost(postid.PostID, &post); err != nil {
		err.Error()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Update success",
			"data": post,
		})
	}
}

func GetImages(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo)
}

func GetImages1(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo1)
}

func GetImages2(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo2)
}

func GetImages3(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo3)
}

func GetImages4(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo4)
}

func GetImages5(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo5)
}

func GetImages6(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo6)
}

func GetImages7(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo7)
}

func GetImages8(c *gin.Context) {
	ID, _ := c.Params.Get("postID")
	postID, _ := strconv.Atoi(ID)
	tempPost := service.GetImages(postID)

	c.Data(200, "image/jpeg", tempPost.Photo8)
}

func SearchPost(c *gin.Context) {
	tag := c.Query("tag")
	//fmt.Println(tag)
	searchpost, err := service.SearchPost(tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"title": "Search Posts",
			"posts": searchpost,
		})
	}
}

func SearchFP(c *gin.Context) {
	userid, _ := c.Params.Get("userid")
	//fmt.Println(tag)
	uid, _ := strconv.Atoi(userid)
	UserFP, err := service.GetMyFP(uid)
	FPlist := []*entity.Post{}
	if err == nil {
		for index, _ := range UserFP {
			FP, err := service.GetPostByID(UserFP[index].PostID)
			if err == nil {
				FPlist = append(FPlist, FP[0])
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"title": "Search Posts",
			"posts": FPlist,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
