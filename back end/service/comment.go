package service

import (
	"bilibili/modles"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func PostComment(ctx *gin.Context)bool{
	var comment1 modles.Comment
	err:=ctx.ShouldBind(&comment1)
	if err!=nil{
		fmt.Println("shouldBind failed",err)
		return false
	}
	id,er:=ctx.Cookie("Id")
	if er!=nil{
		fmt.Println("cookie read failed",er)
		return false
	}
	username:=modles.UserMessage(id).Username
	comment1.Username=username
	comment1.ReleaseTime=time.Now().Format("2006-01-02 15:04:05")
	res:=modles.PostComment(comment1)
	return res
}
func ViewComment(ctx *gin.Context)string{
	bv:=ctx.PostForm("bv")
	comments:=modles.ViesComment(bv)
	comment,err:=disableEscapeHtml(comments)
	if err!=nil{
		fmt.Println("marshal failed",err)
	}
	return comment
}