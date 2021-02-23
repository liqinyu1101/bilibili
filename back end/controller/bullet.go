package controller

import (
	"bilibili/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func ViewBullet(ctx *gin.Context)  {
	res:=service.ViewBullet(ctx)
	if !strings.EqualFold(res,""){
		ctx.JSON(200,res)
	}else {
		ctx.JSON(400,gin.H{
			"message":"failed",
		})
	}
}
func PostBullet(ctx *gin.Context)  {
	res:=service.PostBullet(ctx)
	if !res{
		ctx.JSON(400,gin.H{
			"message":"failed",
		})
	}else {
		ctx.JSON(200,gin.H{
			"message":"successfully",
		})
	}
}
