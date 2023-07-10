package routes

import (
	"Gin_Project/go/controller"
	"Gin_Project/go/entity"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("MySession", store))
	gob.Register(entity.User{})

	r.LoadHTMLGlob("resource/*")
	PostGroup := r.Group("community")
	{
		//帖子首页
		PostGroup.GET("/", controller.PostPage)
		//发表评论
		PostGroup.POST("/create", controller.CreateComment)
		//查看当前帖子的所有评论(点赞数)
		PostGroup.GET("/:id", controller.GetPostComment)
		//查看当前帖子的所有评论(时间)
		PostGroup.GET("/time/:id", controller.TimeSort)
		//点赞某个评论
		PostGroup.PUT("/Like", controller.Like)
		//取消点赞某个评论
		PostGroup.PUT("/Dislike", controller.DisLike)
		//删除某个评论
		PostGroup.DELETE("/:post_id/:comment_id", controller.DeleteCommentById)
		//发表帖子的主页
		PostGroup.GET("/post/create", controller.PostsPage)
		//发表后页面
		PostGroup.POST("/post", controller.CreatePost)
		//查看所有帖子(点赞数)
		PostGroup.GET("/posts", controller.GetPostList)
		//查看所有的帖子(时间)
		PostGroup.GET("/posts/time", controller.GetPosts)
		//输入用户id，搜索该用户的全部帖子
		PostGroup.GET("/search", controller.GetUserID)
		//显示该用户的全部帖子
		PostGroup.GET("/MyPost/:id", controller.GetMyPostList)
		//编辑帖子页面
		PostGroup.GET("/edit/:postid", controller.EditMyPost)
		//提交编辑
		PostGroup.PUT("/edit/:postid/save", controller.UpdatePost)
		//删除帖子
		PostGroup.DELETE("/delete/:postid", controller.DeletePost)
		//获取用户帖子收藏
		PostGroup.GET("/UserFP/:userid", controller.GetUserFP)
		//添加用户收藏
		PostGroup.POST("/AddFP", controller.AddUserFP)
		//取消用户收藏
		PostGroup.DELETE("/DeleteFP", controller.DeleteUserFP)
		//用户点赞帖子
		PostGroup.PUT("/LikePost", controller.LikePost)
		//用户取消点赞帖子
		PostGroup.PUT("/DislikePost", controller.DislikePost)
		//获取图片组：
		PostGroup.GET("/img/:postID", controller.GetImages)
		PostGroup.GET("/img1/:postID", controller.GetImages1)
		PostGroup.GET("/img2/:postID", controller.GetImages2)
		PostGroup.GET("/img3/:postID", controller.GetImages3)
		PostGroup.GET("/img4/:postID", controller.GetImages4)
		PostGroup.GET("/img5/:postID", controller.GetImages5)
		PostGroup.GET("/img6/:postID", controller.GetImages6)
		PostGroup.GET("/img7/:postID", controller.GetImages7)
		PostGroup.GET("/img8/:postID", controller.GetImages8)
		//按照关键词搜索帖子
		PostGroup.GET("/SearchPost", controller.SearchPost)
		//收藏帖子列表
		PostGroup.GET("/FPpost/:userid", controller.SearchFP)
	}

	//景区Spot路由组
	sceneryGroup := r.Group("scenery")
	{
		sceneryGroup.POST("/getcomments", controller.Show_comment)
		sceneryGroup.POST("/getscenery", controller.MainPage)
		sceneryGroup.POST("/getspotpic", controller.Getspotpic)
		sceneryGroup.POST("/postcomment", controller.Make_comment)
		sceneryGroup.POST("/postlikes", controller.Update_likes)
		sceneryGroup.POST("/postorders", controller.Surplus)
		sceneryGroup.POST("/getorders", controller.GetOrders)
		sceneryGroup.POST("/getnickname", controller.Getnickname)
	}

	//景区Spot路由组
	spotGroup := r.Group("spot")
	{
		//增加景点spot
		spotGroup.POST("/spots", controller.CreateSpot)
		//删除某个spot
		spotGroup.DELETE("/spots/:id", controller.DeleteSpotById)
		//查看所有的spot
		spotGroup.GET("/spots", controller.GetSpotList)
		//修改某个spot
		spotGroup.PUT("/spots/:id", controller.UpdateSpot)

		//景区信息录入
		spotGroup.GET("/info", controller.InfoInterface)
		//spotGroup.POST("/info", controller.Info)

		//发送景区数据
		spotGroup.POST("/resource", controller.OrderByScore)
		//发送景区评分排序数据
		spotGroup.POST("/OrderByScore", controller.OrderByScore)
		//发送景区地址筛选数据
		spotGroup.POST("/FilterByAddress", controller.FilterByAddress)
		//发送景区类型筛选数据
		spotGroup.POST("/FilterByType", controller.FilterByType)
		//发送景区名字筛选数据
		spotGroup.POST("/FilterByName", controller.FilterByName)
		//发送景区名字搜索数据
		spotGroup.POST("/SearchByName", controller.SearchByName)
		//发送筛选结果
		spotGroup.POST("/Filter", controller.Filter)
		//发送筛选结果
		spotGroup.POST("/Filter1", controller.Filter1)
		//发送筛选结果
		spotGroup.POST("/search", controller.Search)
		//发送轮播景区
		spotGroup.POST("/display", controller.Display)
	}

	//景区收藏路由组
	spot1Group := r.Group("spot1")
	{
		//增加
		spot1Group.POST("/create", controller.CreateSpot1)
		//删除
		spot1Group.POST("/delete", controller.DeleteSpot1)
		//查看所有
		spot1Group.GET("/spots", controller.GetSpotList1)
	}

	//用户User路由组
	userGroup := r.Group("user")
	{
		//增加用户User
		userGroup.POST("/users", controller.CreateUser)
		//删除某个User
		userGroup.DELETE("/users/:id", controller.DeleteUserById)
		//查看所有的User
		userGroup.GET("/users", controller.GetUserList)
		//修改某个User
		userGroup.PUT("/users/:id", controller.UpdateUser)
		//查找用户名
		userGroup.GET("/username/:id", controller.GetUserName)
		//解析token获取信息
		userGroup.POST("/uid", controller.GetEmail)
		//获取用户头像
		userGroup.GET("/avatar/:userID", controller.GetAvatar)

		//用户注册
		userGroup.POST("/register", controller.Register)
		//用户登录
		userGroup.POST("/login", controller.Login)
		//获取当前用户信息
		userGroup.POST("/resource", controller.GetResource)
		//更新当前用户名字
		userGroup.POST("/updateName", controller.UpdateUserName)

		//获取当前用户收藏列表
		userGroup.POST("/getUserSpot", controller.GetUserSpotList)
		//更新当前用户资料
		userGroup.POST("/update", controller.Update)
		//提供当前用户资料
		userGroup.POST("/getInfo", controller.GetInfo)
		//提供当前用户景区收藏资料
		userGroup.POST("/getCollection", controller.GetCollectSpot)
	}
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Content-Type", "application/json")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
