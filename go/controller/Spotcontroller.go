package controller

import (
	"Gin_Project/go/common/rsp"
	"Gin_Project/go/entity"
	"Gin_Project/go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateSpot(c *gin.Context) {
	//定义一个Spot变量
	var spot entity.Spot
	//将请求中的body数据解析到Spot结构变量中
	if err := c.ShouldBind(&spot); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.CreateSpot(&spot); err != nil {
		rsp.Error(c, "新增景点失败")
	} else {
		rsp.Success(c, "新增景点成功", spot)
	}
}

func GetSpotList(c *gin.Context) {
	if todoList, err := service.GetAllSpot(); err != nil {
		rsp.Error(c, "显示景点列表失败")
	} else {
		rsp.Success(c, "显示景点列表成功", todoList)
	}
}

func OrderByScore(c *gin.Context) {
	token := c.PostForm("token")
	fmt.Println("----------------" + token)
	if token == "null" {
		todoList, err := service.GetSpotByScore()
		for index, value := range todoList {
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		if err != nil {
			rsp.Error(c, "景区按评分排序失败")
			return
		}
		rsp.Success(c, "景区按评分排序成功", todoList)
	}
	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.GetSpotByScore()
	for index, value := range todoList {
		todoList[index].Collection = GetSpot1ById(n, value.Id)
		spotList, _ := service.GetC(value.Id)
		todoList[index].Collections = len(spotList)
	}
	if err != nil {
		rsp.Error(c, "景区按评分排序失败")
		return
	}
	rsp.Success(c, "景区按评分排序成功", todoList)
}

// 显示地址得到的Spot筛选集合
func FilterByAddress(c *gin.Context) {
	address := c.PostForm("address")
	token := c.PostForm("token")
	if token == "null" {
		todoList, err := service.FilterByAddress(address)
		if err != nil {
			rsp.Error(c, "景区按地址筛选失败")
		} else {
			for index, value := range todoList {
				spotList, _ := service.GetC(value.Id)
				todoList[index].Collections = len(spotList)
			}
			rsp.Success(c, "景区按地址筛选成功", todoList)
		}
	}
	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.FilterByAddress(address)
	if err != nil {
		rsp.Error(c, "景区按地址筛选失败")
	} else {
		for index, value := range todoList {
			todoList[index].Collection = GetSpot1ById(n, value.Id)
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		rsp.Success(c, "景区按地址筛选成功", todoList)
	}
}

// 显示类型得到的Spot筛选集合
func FilterByType(c *gin.Context) {
	ttype := c.PostForm("type")
	token := c.PostForm("token")

	if token == "null" {
		todoList, err := service.FilterByType(ttype)
		if err != nil {
			rsp.Error(c, "景区按类型筛选失败")
		} else {
			for index, value := range todoList {
				spotList, _ := service.GetC(value.Id)
				todoList[index].Collections = len(spotList)
			}
			rsp.Success(c, "景区按类型筛选成功", todoList)
		}
	}

	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.FilterByType(ttype)
	if err != nil {
		rsp.Error(c, "景区按类型筛选失败")
	} else {
		for index, value := range todoList {
			todoList[index].Collection = GetSpot1ById(n, value.Id)
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		rsp.Success(c, "景区按类型筛选成功", todoList)
	}
}

// 显示市级地址得到的Spot筛选集合
func FilterByAddress1(c *gin.Context) {
	address1 := c.PostForm("address1")
	token := c.PostForm("token")

	if token == "null" {
		todoList, err := service.FilterByAddress1(address1)
		if err != nil {
			rsp.Error(c, "景区按市级地址筛选失败")
		} else {
			for index, value := range todoList {
				spotList, _ := service.GetC(value.Id)
				todoList[index].Collections = len(spotList)
			}
			rsp.Success(c, "景区按市级地址筛选成功", todoList)
		}
	}

	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.FilterByAddress1(address1)
	if err != nil {
		rsp.Error(c, "景区按市级地址筛选失败")
	} else {
		for index, value := range todoList {
			todoList[index].Collection = GetSpot1ById(n, value.Id)
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		rsp.Success(c, "景区按市级地址筛选成功", todoList)
	}
}

// 显示名字得到的Spot筛选集合
func FilterByName(c *gin.Context) {
	name := c.PostForm("name")
	token := c.PostForm("token")

	if token == "null" {
		todoList, err := service.FilterByName(name)
		if err != nil {
			rsp.Error(c, "景区按名字筛选失败")
		} else {
			for index, value := range todoList {
				spotList, _ := service.GetC(value.Id)
				todoList[index].Collections = len(spotList)
			}
			rsp.Success(c, "景区按名字筛选成功", todoList)
		}
	}

	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.FilterByName(name)
	if err != nil {
		rsp.Error(c, "景区按名字筛选失败")
	} else {
		for index, value := range todoList {
			todoList[index].Collection = GetSpot1ById(n, value.Id)
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		rsp.Success(c, "景区按名字筛选成功", todoList)
	}
}

// 显示地址数组得到的Spot筛选集合
func Filter(c *gin.Context) {
	ttype := c.PostForm("type")
	address := c.PostForm("address")
	token := c.PostForm("token")

	if token == "null" {
		todoList, err := service.Filter(ttype, address)
		if err != nil {
			rsp.Error(c, "景区筛选失败")
		} else {
			for index, value := range todoList {
				spotList, _ := service.GetC(value.Id)
				todoList[index].Collections = len(spotList)
			}
			rsp.Success(c, "景区筛选成功", todoList)
		}
	}

	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.Filter(ttype, address)
	if err != nil {
		rsp.Error(c, "景区筛选失败")
	} else {
		for index, value := range todoList {
			todoList[index].Collection = GetSpot1ById(n, value.Id)
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		rsp.Success(c, "景区筛选成功", todoList)
	}
}

// 显示地址数组得到的Spot筛选集合
func Filter1(c *gin.Context) {
	ttype := c.PostForm("type")
	address := c.PostForm("address")
	token := c.PostForm("token")

	if token == "null" {
		todoList, err := service.Filter1(ttype, address)
		if err != nil {
			rsp.Error(c, "景区筛选失败")
		} else {
			for index, value := range todoList {
				spotList, _ := service.GetC(value.Id)
				todoList[index].Collections = len(spotList)
			}
			rsp.Success(c, "景区筛选成功", todoList)
		}
	}

	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.Filter1(ttype, address)
	if err != nil {
		rsp.Error(c, "景区筛选失败")
	} else {
		for index, value := range todoList {
			todoList[index].Collection = GetSpot1ById(n, value.Id)
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		rsp.Success(c, "景区筛选成功", todoList)
	}
}

// 显示名字得到的Spot模糊搜索集合
func SearchByName(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	name := c.PostForm("name")
	token := c.PostForm("token")

	if token == "null" {
		todoList, err := service.SearchByName(name)
		if err != nil {
			rsp.Error(c, "景区按名字搜索失败")
		} else {
			for index, value := range todoList {
				spotList, _ := service.GetC(value.Id)
				todoList[index].Collections = len(spotList)
			}
			rsp.Success(c, "景区按名字搜索成功", todoList)
		}
	}

	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.SearchByName(name)
	if err != nil {
		rsp.Error(c, "景区按名字搜索失败")
	} else {
		for index, value := range todoList {
			todoList[index].Collection = GetSpot1ById(n, value.Id)
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		rsp.Success(c, "景区按名字搜索成功", todoList)
	}
}

func Search(c *gin.Context) {
	name := c.PostForm("name")
	token := c.PostForm("token")

	if token == "null" {
		todoList, err := service.Search(name)
		if err != nil {
			rsp.Error(c, "景区搜索失败")
		} else {
			for index, value := range todoList {
				spotList, _ := service.GetC(value.Id)
				todoList[index].Collections = len(spotList)
			}
			rsp.Success(c, "景区搜索成功", todoList)
		}
	}

	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	todoList, err := service.Search(name)
	if err != nil {
		rsp.Error(c, "景区搜索失败")
	} else {
		for index, value := range todoList {
			todoList[index].Collection = GetSpot1ById(n, value.Id)
			spotList, _ := service.GetC(value.Id)
			todoList[index].Collections = len(spotList)
		}
		rsp.Success(c, "景区搜索成功", todoList)
	}
}

func UpdateSp(c *gin.Context) {
	sid := c.PostForm("sid")
	score := c.PostForm("score")

	n, _ := strconv.Atoi(sid)

	spot, _ := service.GetSpotById(n)

	n1, _ := strconv.ParseFloat(score, 32)
	spot.Score = float32(n1)

	if err := service.UpdateSpot(&spot); err != nil {
		rsp.Error(c, "更新景点信息失败")
		return
	} else {
		rsp.Success(c, "更新景点信息成功", spot)
	}
}

func UpdateSpot(c *gin.Context) {
	//得到URL上的id信息
	token := c.PostForm("token")
	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)

	i, err := strconv.Atoi(string(user.Id))
	if err != nil {
		rsp.Error(c, "景点更新，景点ID不合法")
		return
	}
	var spot entity.Spot
	spot, err = service.GetSpotById(i)
	if err != nil {
		rsp.Error(c, "景点更新，景点不存在")
		return
	}

	//将请求中的body数据解析到User结构变量中
	if err := c.ShouldBind(&spot); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := service.UpdateSpot(&spot); err != nil {
		rsp.Error(c, "更新景点信息失败")
		return
	} else {
		rsp.Success(c, "更新景点信息成功", spot)
	}
}

func DeleteSpotById(c *gin.Context) {
	token := c.PostForm("token")
	var claim *entity.Claims
	claim, _ = service.ParseTokenHs256(c, token)
	email := claim.Email
	user, _ := service.GetUserByEmail(email)
	n := user.Id

	//将id传给service层的UpdateUser方法，进行User的删除
	if err := service.DeleteSpotById(string(n)); err != nil {
		rsp.Error(c, "删除景点失败")
	} else {
		rsp.Success(c, "删除景点成功", string(n))
	}
}

// 景区信息录入接口
func InfoInterface(c *gin.Context) {
	c.HTML(http.StatusOK, "spotinfo.html", nil)
}

func Display(c *gin.Context) {

	spotList1 := []*entity.Spot{}
	todoList, err := service.GetSpotByScore()
	for index, value := range todoList {
		spotList, _ := service.GetC(value.Id)
		value.Collections = len(spotList)
		if index < 6 {
			spotList1 = append(spotList1, value)
		}
	}
	if err != nil {
		rsp.Error(c, "景区按评分排序失败")
		return
	}
	rsp.Success(c, "景区按评分排序成功", spotList1)
}

// 景区信息录入实现
//func Info(c *gin.Context) {
//	//将请求中的body数据解析到User结构变量中
//	var reqUser entity.Spot
//	if err := c.ShouldBind(&reqUser); err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//		return
//	}
//	//获取景区信息
//	name := reqUser.Name
//	address := reqUser.Address
//	text := reqUser.Text
//
//	newUser := &entity.Spot{
//		Name:    name,
//		Address: address,
//	}
//
//	//创建该景区
//	if err := service.CreateSpot(newUser); err != nil {
//		rsp.Error(c, "景区信息录入，失败")
//		return
//	}
//	id := newUser.Id
//
//	//创建文本文件存储描述信息
//	filePath := "D:/Code/GoProject/GoLand/src/gin/gin_demo/file/spot/spot_text/"
//	filePath += strconv.Itoa(id) + ".txt"
//
//	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil {
//		rsp.Error(c, "景区信息录入，描述文件打开失败")
//		return
//	}
//	defer func(file *os.File) {
//		err := file.Close()
//		if err != nil {
//			rsp.Error(c, "景区信息录入，描述文件关闭失败")
//			return
//		}
//	}(file)
//
//	writer := bufio.NewWriter(file)
//	_, err = writer.WriteString(text)
//	if err != nil {
//		rsp.Error(c, "景区信息录入，文件输入失败")
//		return
//	}
//	err = writer.Flush()
//	if err != nil {
//		rsp.Error(c, "景区信息录入,文本输入失败")
//		return
//	}
//	newUser.Text = strconv.Itoa(id) + ".txt"
//
//	//从请求中读取图片文件
//	f, err := c.FormFile("picture")
//	if err != nil {
//		rsp.Error(c, "景区信息录入，景图读取失败")
//		return
//	} else {
//		//将文件保存在服务器
//		dst := path.Join("./file./spot/spot_image/", strconv.Itoa(id)+".jpg")
//
//		if err := c.SaveUploadedFile(f, dst); err != nil {
//			rsp.Error(c, "景区信息录入，景图保存服务器错误")
//			return
//		}
//
//		newUser.Picture = strconv.Itoa(id) + ".jpg"
//	}
//	rsp.Success(c, "景区信息录入，成功", newUser)
//}
