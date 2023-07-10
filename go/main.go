package main

import (
	"Gin_Project/go/dao"
	"Gin_Project/go/entity"
	"Gin_Project/go/routes"
)

func main() {
	//连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.Close()
	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.Comment{})
	dao.UserSession.AutoMigrate(&entity.User{})
	dao.PostSession.AutoMigrate(&entity.Post{})
	dao.UserFP.AutoMigrate(&entity.UserFP{})
	//注册路由
	r := routes.SetRouter()
	//启动端口为8080的项目
	r.Run(":8080")
}
