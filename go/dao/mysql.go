package dao

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"post"`
}

const DRIVER = "mysql"

var SqlSession *gorm.DB  //评论
var UserSession *gorm.DB //用户
var PostSession *gorm.DB //帖子
var UserFP *gorm.DB      //收藏

func (c *conf) getConf() *conf {
	//读取resources/application.yaml文件
	yamlFile, err := ioutil.ReadFile("resource/application.yaml")
	//若出现错误，打印错误提示
	if err != nil {
		fmt.Println(err.Error())
	}
	//将读取的字符串转换成结构体conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func InitMySql() (err error) {
	var c conf
	//获取yaml配置参数
	conf := c.getConf()
	//将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	//连接数据库
	SqlSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	UserSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	PostSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	UserFP, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	//验证数据库连接是否成功，若成功，则无异常
	e1 := SqlSession.DB().Ping()
	e2 := UserSession.DB().Ping()
	e3 := PostSession.DB().Ping()
	e4 := UserFP.DB().Ping()
	if e1 == nil && e2 == nil && e3 == nil && e4 == nil {
		return nil
	} else if e1 != nil {
		return e1
	} else if e2 != nil {
		return e2
	} else if e3 != nil {
		return e3
	} else {
		return e4
	}
}

func Close() {
	err := SqlSession.Close()
	if err != nil {
		return
	}
	err = UserSession.Close()
	if err != nil {
		return
	}
	err = PostSession.Close()
	if err != nil {
		return
	}
	err = UserFP.Close()
	if err != nil {
		return
	}
}

func GetDb() (db *sql.DB, err error) {
	config := mysql.Config{
		User:                 "root",
		Passwd:               "dxy156155",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "community",
		AllowNativePasswords: true,
	}
	db, err = sql.Open("mysql", config.FormatDSN())
	return db, err
}
