package controller

import (
	"bilibili/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func UploadVideo(ctx *gin.Context)  {
	res:=service.UploadVideo(ctx)
	if res{
		ctx.JSON(200,gin.H{
			"message":"successfully",
		})
	}else {
		ctx.JSON(400,gin.H{
			"message":"failed",
		})
	}
}
func RecommendVideo(ctx *gin.Context){
	videos:=service.RecommendVideo(ctx)
	fmt.Println(videos)
	if videos==""{
		ctx.JSON(400,gin.H{
			"message":"something wrong",
		})
	}else {
		ctx.JSON(200,videos)
	}
}
//1成功，0未知错误，-1已经投过币，-2硬币不足
func InsertCoin(ctx *gin.Context)  {
	res:=service.InsertCoin(ctx)
	if res==1{
		ctx.JSON(200,gin.H{
			"message":"successfully",
		})
	}else if res==0{
		ctx.JSON(400,gin.H{
			"message":"something wrong",
		})
	}else if res==-1{
		ctx.JSON(200,gin.H{
			"message":"have inserted coin",
		})
	}else {
		ctx.JSON(200,gin.H{
			"message":"no enough coin",
		})
	}
}
//1 已经点赞 0未知错误 -1 取消点赞
func ThumbsUp(ctx *gin.Context){
	res:=service.ThumbsUp(ctx)
	if res==1{
		ctx.JSON(200,gin.H{
			"message":"thumbsUp successful",
		})
	}else if res==0{
		ctx.JSON(400,gin.H{
			"message":"something wrong",
		})
	}else {
		ctx.JSON(200,gin.H{
			"message":"cancel thumbsUp",
		})
	}
}

func SearchVideo(ctx *gin.Context){
	res:=service.SearchVideo(ctx)
	if !strings.EqualFold("",res){
		ctx.JSON(200,res)
	}else {
		ctx.JSON(400,gin.H{
			"message":"Search failed",
		})
	}
}
//1 已经收藏 0未知错误 -1 取消收藏
func Collect(ctx *gin.Context){
	res:=service.Collect(ctx)
	if res==1{
		ctx.JSON(200,gin.H{
			"message":"Collect successful",
		})
	}else if res==0{
		ctx.JSON(400,gin.H{
			"message":"something wrong",
		})
	}else {
		ctx.JSON(200,gin.H{
			"message":"cancel Collect",
		})
	}
}
func Video(ctx *gin.Context){
	ctx.HTML(200,"",gin.H{})
}
func Links(ctx *gin.Context){
	res:=service.Links(ctx)
	if res{
		ctx.JSON(200,gin.H{
			"message":"successfully",
		})
	}else {
		ctx.JSON(400,gin.H{
			"message":"failed",
		})
	}
}
func ViewPostVideo(ctx *gin.Context){
	res:=service.ViewPostVideo(ctx)
	ctx.JSON(200,res)
}
func VideoTimesAdd(ctx *gin.Context) {
	res:=service.VideoTimesAdd(ctx)
	if res{
		ctx.JSON(200,gin.H{
			"message":"successfully",
		})
	}else {
		ctx.JSON(400,gin.H{
			"message":"failed",
		})
	}
}
func DeleteVideo(ctx *gin.Context){
	res:=service.DeleteVideo(ctx)
	if res{
		ctx.JSON(200,gin.H{
			"message":"successfully",
		})
	}else {
		ctx.JSON(400,gin.H{
			"message":"failed",
		})
	}
}
