package service

import (
	"bilibili/modles"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ViewBullet(ctx *gin.Context) string{
	bv:=ctx.PostForm("bv")
	bullets:=modles.ViewBullet(bv)
	bullet,err:=disableEscapeHtml(bullets)
	if err!=nil{
		fmt.Println("marshal wrong",err)
		return ""
	}
	return bullet
}
func PostBullet(ctx *gin.Context)bool{
	var content modles.Bullet
	err:=ctx.ShouldBind(&content)
	if err!=nil{
		fmt.Println("shouldBind failed",err)
	}
	res:=modles.PostBullet(content)
	return res
}