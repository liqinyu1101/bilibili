package controller

import "github.com/gin-gonic/gin"

func CookieValue(ctx *gin.Context)  {
	cookieValue,_:=ctx.Cookie("Id")
	ctx.JSON(200,gin.H{
		"cookieValue":cookieValue,
	})
}