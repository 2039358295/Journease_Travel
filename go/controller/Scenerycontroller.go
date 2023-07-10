package controller

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
	"Gin_Project/go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var _id int
var _name string
var _price1 float32
var _price2 float32
var _price3 float32
var _address string
var _introduction string
var _score float32
var _userid int
var _num int
var photo1 string
var photo2 string
var photo3 string
var photo4 string
var photo5 string
var photo6 string
var comment [6]string
var new_comment entity.Comments

func GetID(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	type i struct {
		UserID    int `json:"userid"`
		SceneryID int `json:"sceneryid"`
	}

	var _i i
	c.BindJSON(&_i)
	_id = _i.SceneryID
	_userid = _i.UserID
}

func Getnickname(c *gin.Context) {
	GetID(c)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT nick_name,avatar FROM sys_user where id=?", _userid)

	var _nickname string
	var _avatar string
	rows.Next()
	err = rows.Scan(&_nickname, &_avatar)
	//rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"nickname": _nickname,
		"avatar":   _avatar,
	})
}

func Update_likes(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	type t struct {
		ID    int `json:"id"`
		Likes int `json:"likes"`
	}

	var _t t
	c.BindJSON(&_t)
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ret, err := db.Exec("update comments set likes=? where id=?", _t.Likes, _t.ID)
	if err != nil {
		fmt.Println("insert err = ", err, ret)
		return
	}
}

func Show_comment(c *gin.Context) {
	GetID(c)
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	commentList, err := service.GetAllComments(_id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"commentList": commentList,
	})
}

func Get_photos(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	photosList, err := service.GetAllPhotos()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"spotpic": photosList,
	})
}

func Make_comment(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	c.BindJSON(&new_comment)
	fmt.Println(new_comment)

	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ret, err := db.Exec("insert into comments (comment,userid,sceneryid,likes,time,avatar,name,rating) values (?,?,?,0,?,?,?,?)", new_comment.Comment, new_comment.Userid, _id, new_comment.Time, new_comment.Avatar, new_comment.Name, new_comment.Rating)
	if err != nil {
		fmt.Println("insert err = ", err, ret)
		return
	}
	Score(c)
}

func Score(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")

	_num++
	my_score := new_comment.Rating
	_score = (_score*float32(_num-1) + float32(my_score)) / float32(_num)
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	cache, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", _score), 64)
	_score = float32(cache)
	ret, err := db.Exec("update scenery set score=?,num=? where id=?", _score, _num, _id)

	//spot, _ := service.GetSpotById(2)
	//	//spot.Score = 0
	//	//service.UpdateSpot(&spot)

	if err != nil {
		fmt.Println("insert err = ", err, ret)
		return
	}

}

func Surplus(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	var order entity.Userorders
	c.BindJSON(&order)
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ret, err := db.Exec("insert into userorders (userid,date,sceneryid,type,num,sceneryname,price) values (?,?,?,?,?,?,?)", _userid, order.Date, _id, order.Type, order.Num, order.Sceneryname, order.Price)
	if err != nil {
		fmt.Println("insert err = ", err, ret)
		return
	}
}

func Getspotpic_son(c *gin.Context) {
	i := 0
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT comment FROM comments where sceneryid=? order by likes desc", _id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&comment[i])
		//rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		if i == 5 {
			break
		}
		//err = rows.Err()
		//if err != nil {
		//	log.Fatal(err)
		//}
		i++
	}
	db.Close()
}

func Getspotpic(c *gin.Context) {
	GetID(c)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	db1, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db1.Query("SELECT * FROM scenery")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var goal_id int

	for rows.Next() {
		err = rows.Scan(&goal_id, &_name, &_price1, &_price2, &_price3, &_address, &_introduction, &_score, &_num, &photo1, &photo2, &photo3, &photo4, &photo5, &photo6)
		//rows.Err()
		if _id == goal_id {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//err = rows.Err()
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
	db1.Close()
	Getspotpic_son(c)
	_db, _err := dao.GetDb()
	if _err != nil {
		log.Fatal(err)
	}
	defer _db.Close()
	if _id == goal_id {
		ret1, err1 := _db.Exec("insert into spotpics (id,photo,comment) values (?,?,?)", _id, photo1, comment[0])
		if err1 != nil {
			fmt.Println("insert err1 = ", err1, ret1)
			return
		}
		ret2, err2 := _db.Exec("insert into spotpics (id,photo,comment) values (?,?,?)", _id, photo2, comment[1])
		if err2 != nil {
			fmt.Println("insert err2 = ", err2, ret2)
			return
		}
		ret3, err3 := _db.Exec("insert into spotpics (id,photo,comment) values (?,?,?)", _id, photo3, comment[2])
		if err3 != nil {
			fmt.Println("insert err3 = ", err3, ret3)
			return
		}
		ret4, err4 := _db.Exec("insert into spotpics (id,photo,comment) values (?,?,?)", _id, photo4, comment[3])
		if err4 != nil {
			fmt.Println("insert err4 = ", err4, ret4)
			return
		}
		ret5, err5 := _db.Exec("insert into spotpics (id,photo,comment) values (?,?,?)", _id, photo5, comment[4])
		if err5 != nil {
			fmt.Println("insert err5 = ", err5, ret5)
			return
		}
		ret6, err6 := _db.Exec("insert into spotpics (id,photo,comment) values (?,?,?)", _id, photo6, comment[5])
		if err6 != nil {
			fmt.Println("insert err6 = ", err6, ret6)
			return
		}
		Get_photos(c)
		_ret, _err := _db.Exec("delete from spotpics")
		if _err != nil {
			fmt.Println("insert _err = ", _err, _ret)
			return
		}
	}

}

func MainPage(c *gin.Context) {
	GetID(c)
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM scenery")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var goal_id int

	for rows.Next() {
		err = rows.Scan(&goal_id, &_name, &_price1, &_price2, &_price3, &_address, &_introduction, &_score, &_num, &photo1, &photo2, &photo3, &photo4, &photo5, &photo6)
		//rows.Err()
		if _id == goal_id {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//err = rows.Err()
		//if err != nil {
		//	log.Fatal(err)
		//}
	}
	if _id == goal_id {
		c.JSON(200, gin.H{
			"id":           _id,
			"name":         _name,
			"price1":       _price1,
			"price2":       _price2,
			"price3":       _price3,
			"address":      _address,
			"introduction": _introduction,
			"score":        _score,
			"photo1":       photo1,
			"photo2":       photo2,
			"photo3":       photo3,
			"photo4":       photo4,
			"photo5":       photo5,
			"photo6":       photo6,
		})
	} else {
		c.String(200, "error")
	}
}

func GetOrders(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json")
	type i struct {
		UserID int `json:"userid"`
	}

	var _i i
	c.BindJSON(&_i)
	db, err := dao.GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	orderList, err := service.GetAllOrders(_i.UserID)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"orderList": orderList,
	})
}
