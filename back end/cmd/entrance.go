package cmd

import (
	"bilibili/controller"
	"bilibili/middleware"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "static")
	//读取cookie的value
	router.GET("/cookie",middleware.Cookie(),controller.CookieValue)
	routerBilibili:=router.Group("/bilibili")
	//打开注册界面
	routerBilibili.GET("/register", controller.Register)
	//发送phone,username,password
	routerBilibili.POST("/register", controller.Register_)
	//发送username 可用的话返回1 不可以返回0  就知道了
	routerBilibili.POST("/register/username",controller.RegisterUsername)
	//发送phone 同上
	routerBilibili.POST("/register/phone",controller.RegisterPhone)
	//打开登录界面
	routerBilibili.GET("/login", controller.Login)
	//发送phone或者username 和password
	routerBilibili.POST("/login", controller.Login_)
	//拿到登录者的信息
	routerBilibili.GET("/user", middleware.AddCoin(),controller.SelfMessage)
	//打开上传头像的界面
	routerBilibili.GET("/face/upload", middleware.Cookie(),  controller.UploadFace_)
	//发送文件face
	routerBilibili.POST("/face/upload", middleware.Cookie(),controller.UploadFace)
	//登出 有问题
	routerBilibili.GET("/login/logout",middleware.Cookie(),controller.Logout)
	//点关注或者取消关注 up（这个是被关注人的id）
	routerBilibili.POST("/follow",middleware.Cookie(),controller.Follow)
	//搜索keyword  query参数
	routerBilibili.GET("/user/search",controller.SearchUser)
	//其他人的用户信息 发送id
	routerBilibili.GET("/user/other",controller.UserMessage)
	//看一个人的点赞数和播放次数 发送id
	routerBilibili.GET("/ThumbsUpAndTimes",controller.ThumbsUpAndTimes)
	videoRouter:=routerBilibili.Group("/video")
	//发送 bv  看登录的人对视频是否关注，投币，点赞
	videoRouter.GET("/status",controller.VideoStatus)
	//上传视频 title commit tag video cover
	videoRouter.POST("/upload",middleware.Cookie(),controller.UploadVideo)
	//查找视频 发送keyword
	videoRouter.GET("/search",controller.SearchVideo)
	//发送tag和number 如果没有tag 就观看次数推荐number个视频
	videoRouter.POST("/recommend",controller.RecommendVideo)
	//投币 发送bv
	videoRouter.POST("/coin",middleware.Cookie(),controller.InsertCoin)
	//点赞或取消点赞 发送bv
	videoRouter.POST("/thumbsUp",middleware.Cookie(),controller.ThumbsUp)
	//收藏火取消收藏 发送bv
	videoRouter.POST("/collect",middleware.Cookie(),controller.Collect)
	//打开视频页
	videoRouter.GET("",controller.Video)
	//一键三连
	videoRouter.POST("/links",middleware.Cookie(),controller.Links)
	//发送id 看发布的视频
	videoRouter.GET("/videos",controller.ViewPostVideo)
	//增加视频的观看次数
	videoRouter.POST("/times/add",controller.VideoTimesAdd)
	//删除视频，发送bv
	videoRouter.POST("/delete",middleware.Cookie(),controller.DeleteVideo)
	commentRouter:=routerBilibili.Group("/comment")
	//发送content bv 发送评论
	commentRouter.POST("/post",middleware.Cookie(),controller.PostComment)
	//发送bv 看视频的评论
	commentRouter.GET("/view",controller.ViewComment)
	bulletRouter:=routerBilibili.Group("/bullet")
	//发送bv 看弹幕
	bulletRouter.GET("/view",controller.ViewBullet)
	//发送time bv content
	bulletRouter.POST("/post",middleware.Cookie(),controller.PostBullet)
	router.Run(":80")
}
