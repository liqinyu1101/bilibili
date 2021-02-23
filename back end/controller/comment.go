package controller

import (
	"bilibili/service"
	"github.com/gin-gonic/gin"
)

func PostComment(ctx *gin.Context) {
	res:=service.PostComment(ctx)
	if res {
		ctx.JSON(200,gin.H{
			"message":"post successfully",
		})
	}else {
		ctx.JSON(200,gin.H{
			"message":"post failed",
		})
	}
}
func ViewComment(ctx *gin.Context){
	res:=service.ViewComment(ctx)
	ctx.JSON(200,res)
}