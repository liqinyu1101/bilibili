package controller

import (
	"bilibili/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Register_(ctx *gin.Context) {
	res := service.Register_(ctx)
	if res == 0 {
		fmt.Println("register failed，username or phone has benn used")
		ctx.JSON(200, gin.H{
			"message": "register failed，username or phone has benn used",
		})
	} else {
		fmt.Printf("账号id=%d", res)
		ctx.JSON(200, gin.H{
			"message": "register successfully,your id =" + strconv.Itoa(res),
		})
	}

}
func Register(ctx *gin.Context) {
	ctx.HTML(200, "register.html", gin.H{})
}
func Login(ctx *gin.Context) {
	ctx.HTML(200, "test.html", gin.H{})
}
func Login_(ctx *gin.Context) {
	res := service.Login_(ctx)
	if !res {
		ctx.JSON(400, gin.H{
			"message": "login failed and please check it out",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "login successful",
		})
	}
}
func SelfMessage(ctx *gin.Context) {
	res := service.SelfMessage(ctx)
	ctx.JSON(200, res)
}
func UploadFace(ctx *gin.Context) {
	res := service.UploadFace(ctx)
	if !res {
		ctx.JSON(400, gin.H{
			"message": "upload failed",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "successfully",
		})
	}
}
func UploadFace_(ctx *gin.Context) {
	ctx.HTML(200, "upload.html", nil)
}
func Logout(ctx *gin.Context) {
	cookie,err:=ctx.Request.Cookie("Id")
	if err!=nil{
		ctx.JSON(400,gin.H{
			"message":"get cookie failed",
		})
	}else {
		cookie.MaxAge=-1
		//cookie.Expires=time.Now().Add(-time.Hour)
		ctx.JSON(200,gin.H{
			"message":"logout successful",
		})
	}
}
func Follow(ctx *gin.Context)  {
	res:=service.Follow(ctx)
	if res==1{
		ctx.JSON(200,gin.H{
			"message":" Cancel attention successfully",
		})
	}else if res==-1{
		ctx.JSON(200,gin.H{
			"message":" Follow successfully",
		})
	}else {
		ctx.JSON(400,gin.H{
			"message":"something wrong",
		})
	}
}
func SearchUser(ctx *gin.Context){
	res:=service.SearchUser(ctx)
	if !strings.EqualFold("",res){
		ctx.JSON(200,res)
	}else {
		ctx.JSON(400,gin.H{
			"message":"Search failed",
		})
	}
}
func UserMessage(ctx *gin.Context){
	res:=service.UserMessage(ctx)
	ctx.JSON(200,res)
}
func VideoStatus(ctx *gin.Context){
	res:=service.VideoStatus(ctx)
	ctx.JSON(200,res)
}
func ThumbsUpAndTimes(ctx *gin.Context){
	res:=service.ThumbsUpAndTimes(ctx)
	ctx.JSON(200,res)
}
func RegisterUsername(ctx *gin.Context){
	res:=service.RegisterUsername(ctx)
	if res {
		ctx.JSON(200,gin.H{
			"flag":1,
		})
	}else {
		ctx.JSON(200,gin.H{
			"flag":0,
		})
	}
}
func RegisterPhone(ctx *gin.Context){
	res:=service.RegisterPhone(ctx)
	if res {
		ctx.JSON(200,gin.H{
			"flag":1,
		})
	}else {
		ctx.JSON(200,gin.H{
			"flag":0,
		})
	}
}
