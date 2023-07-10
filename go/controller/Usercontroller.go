package controller

import (
	"Gin_Project/go/common/rsp"
	"Gin_Project/go/entity"
	"Gin_Project/go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	//定义一个User变量
	var user entity.User
	//将请求中的body数据解析到User结构变量中
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.CreateUser(&user); err != nil {
		rsp.Error(c, "新增用户失败")
	} else {
		rsp.Success(c, "新增用户成功", user)
	}
}

func GetUserList(c *gin.Context) {
	if todoList, err := service.GetAllUser(); err != nil {
		rsp.Error(c, "显示用户列表失败")
	} else {
		rsp.Success(c, "显示用户列表成功", todoList)
	}
}

func UpdateUser(c *gin.Context) {
	//得到URL上的id信息
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		rsp.Error(c, "用户更新，用户ID不合法")
		return
	}
	var user entity.User
	user, err = service.GetUserById(i)
	if err != nil {
		rsp.Error(c, "用户更新，用户不存在")
		return
	}

	//将请求中的body数据解析到User结构变量中
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.UpdateUser(&user); err != nil {
		rsp.Error(c, "更新用户信息失败")
		return
	} else {
		rsp.Success(c, "更新用户信息成功", user)
	}
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	//将id传给service层的UpdateUser方法，进行User的删除
	if err := service.DeleteUserById(id); err != nil {
		rsp.Error(c, "删除用户失败")
	} else {
		rsp.Success(c, "删除用户成功", id)
	}
}

// 用户注册接口
func Register(c *gin.Context) {
	//获取注册信息
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	email := c.PostForm("email")
	password1 := c.PostForm("password1")

	if nickname == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "昵称为空",
		})
		return
	}

	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "邮箱为空",
		})
		return
	}

	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "密码为空",
		})
		return
	}

	if password1 == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "确认密码为空",
		})
		return
	}

	if password1 != password {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "密码不一致",
		})
		return
	}

	//邮箱查重
	_, err := service.GetUserByEmail(email)
	if err == nil {
		rsp.Error(c, "邮箱已存在")
		return
	}

	//对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		rsp.Error(c, "密码出错")
		return
	}

	filePath := "D:\\GoProject\\GoLand\\src\\Gin_Project\\vue_pro\\vue_pro\\src\\assets\\logo.jpg"
	imageData, err := ioutil.ReadFile(filePath)

	token, _ := service.GenerateToken(email, password)
	newUser := &entity.User{
		NickName: nickname,
		Password: string(hashPassword),
		Email:    email,
		Token:    token,
		Avatar:   imageData,
	}
	fmt.Println("---------" + newUser.NickName)

	//创建该用户
	if err := service.CreateUser(newUser); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "服务器出错",
			"data": newUser,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": token,
	})
}

// 用户登录接口
func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	/*
		var reqUser entity.User
		//将请求中的body数据解析到User结构变量中
		if err := c.ShouldBind(&reqUser); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		email := reqUser.Email
		password := reqUser.Password
	*/

	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "邮箱为空",
		})
		return
	}

	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "密码为空",
		})
		return
	}

	//查询用户
	var user entity.User
	user, err := service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "邮箱不存在",
		})
		return
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}

	token, _ := service.GenerateToken(email, password)
	user.Token = token

	//创建该用户
	if err := service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "服务器出错",
			"data": user,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "登录成功",
		"ava":      user.Avatar,
		"userName": user.Name,
		"token":    user.Token,
		"uid":      user.Id,
	})
}

// 用户资料接口
func GetInfo(c *gin.Context) {
	uid := c.PostForm("uid")
	n, _ := strconv.Atoi(uid)
	//查询用户
	var user entity.User
	user, err := service.GetUserById(n)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "用户错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "提交信息成功",
		"ava":      user.Avatar,
		"nickname": user.NickName,
		"email":    user.Email,
	})
}

// 用户景区收藏接口
func GetCollectSpot(c *gin.Context) {
	uid := c.PostForm("uid")
	n, _ := strconv.Atoi(uid)

	todoList, err := service.GetUserSpot(n)
	spotList := []*entity.Spot{}

	if err != nil {
		rsp.Error(c, "个人景点收藏提供失败")
	} else {
		for index, value := range todoList {
			spot, err := service.GetSpotById(todoList[index].Sid)
			if err == nil {
				spot.Collection = 1
				spot1List, _ := service.GetC(spot.Id)
				spot.Collections = len(spot1List)

				spotList = append(spotList, &spot)
				fmt.Println(value)
			}
		}
		rsp.Success(c, "个人景点收藏提供成功", spotList)
	}
}

func GetResource(c *gin.Context) {
	token := c.PostForm("token")

	var claim *entity.Claims
	claim, err := service.ParseTokenHs256(c, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "token失效",
		})
		return
	}
	email := claim.Email

	//查询用户
	var user entity.User
	user, err = service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户不存在",
		})
		return
	}

	res := []entity.Res{
		{1, "index", "首页", "/index", "index"},
		{2, "spot", "景区", "/spot", "spot"},
		{3, "community", "社区", "/community", "community"},
		{4, "center", "个人中心", "/center", "center"},
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "信息请求成功",
		"data":     res,
		"identity": "user",
		"email":    user.Email,
	})
}

// 更新昵称
func UpdateUserName(c *gin.Context) {
	token := c.PostForm("token")
	var claim *entity.Claims
	claim, err := service.ParseTokenHs256(c, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "token失效",
		})
		return
	}
	email := claim.Email

	//查询用户
	var user entity.User
	user, err = service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户不存在",
		})
		return
	}

	user.Name = c.PostForm("name")

	if err := service.UpdateUser(&user); err != nil {
		rsp.Error(c, "更新用户信息失败")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "更新用户信息成功",
		})
	}
}

// 更新密码
func UpdateUserPwd(c *gin.Context) {
	token := c.PostForm("token")
	var claim *entity.Claims
	claim, err := service.ParseTokenHs256(c, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "token失效",
		})
		return
	}
	email := claim.Email

	//查询用户
	var user entity.User
	user, err = service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户不存在",
		})
		return
	}

	password := c.PostForm("password")
	//然后对密码进行加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		rsp.Error(c, "用户密码加密错误")
		return
	}

	user.Password = string(hashPassword)

	if err := service.UpdateUser(&user); err != nil {
		rsp.Error(c, "更新用户信息失败")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "更新用户信息成功",
		})
	}
}

// 更新邮箱
func UpdateUserEmail(c *gin.Context) {
	token := c.PostForm("token")
	var claim *entity.Claims
	claim, err := service.ParseTokenHs256(c, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "token失效",
		})
		return
	}

	//判断是否重复
	email1 := c.PostForm("email")
	_, err = service.GetUserByEmail(email1)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "邮箱已重复",
		})
		return
	}

	email := claim.Email
	var user entity.User
	user, err = service.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户不存在",
		})
		return
	}

	token, _ = service.GenerateToken(email, user.Password)
	user.Email = email1
	user.Token = token

	if err := service.UpdateUser(&user); err != nil {
		rsp.Error(c, "更新用户信息失败")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "更新用户信息成功",
		})
	}
}

// 更新用户资料
func Update(c *gin.Context) {
	uid := c.PostForm("uid")
	email := c.PostForm("email")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")

	fileHeader := c.Request.MultipartForm.File["avatar[raw]"]
	file, err := fileHeader[0].Open()
	defer file.Close()
	if err != nil {
		return
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	avatar := fileData

	n, _ := strconv.Atoi(uid)
	//查询用户
	var user entity.User
	user, err = service.GetUserById(n)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}

	//判断是否重复
	if email != "" && user.Email != email {
		_, err = service.GetUserByEmail(email)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "邮箱已重复",
			})
			return
		}
		user.Email = email
	}

	//录入密码
	if password != "" {
		//然后对密码进行加密
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "密码更新失败",
			})
			return
		}
		user.Password = string(hashPassword)
	}

	//录入昵称
	if nickname != "" {
		user.NickName = nickname
	}

	if avatar != nil {
		user.Avatar = avatar
	}

	if err := service.UpdateUser(&user); err != nil {
		rsp.Error(c, "服务器出错")
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "更新成功",
		})
	}
	return
}

func GetUserSpotList(c *gin.Context) {
	todoList, err := service.GetUserSpot(811)
	if err != nil {
		rsp.Error(c, "显示景点收藏列表失败")
	}

	spotList := make([]*entity.Spot, 0)
	for index := range todoList {
		spot, err := service.GetSpotById(todoList[index].Sid)
		if err != nil {
			rsp.Error(c, "景点获取失败")
		}
		spotList = append(spotList, &spot)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "用户景点收藏列表",
		"data": spotList,
	})
}

func GetUserName(c *gin.Context) {
	ID := c.Param("id")
	userID, err := strconv.Atoi(ID)
	if err != nil {
		return
	}
	user := service.CheckUserName(userID)
	if user.NickName != "" {
		c.JSON(200, gin.H{
			"name": user.NickName,
		})
	} else {
		c.JSON(400, gin.H{
			"name": "null",
		})
	}
}

func GetEmail(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	var token string
	err := c.BindJSON(&token)
	var claim *entity.Claims
	claim, err = service.ParseTokenHs256(c, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "token失效",
		})
		return
	}
	email := claim.Email
	uid, err := service.GetUserByEmail(email)
	c.JSON(200, gin.H{
		"email": email,
		"uid":   uid.Id,
	})
}

func GetAvatar(c *gin.Context) {
	ID, _ := c.Params.Get("userID")
	userID, _ := strconv.Atoi(ID)
	temp := service.GetUserAvatar(userID)

	c.Data(200, "image/jpeg", temp.Avatar)
}
