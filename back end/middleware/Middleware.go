package middleware

import (
	"bilibili/modles"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
func Cookie() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, err := c.Request.Cookie("Id")
		//fmt.Println(c1.MaxAge,c1.Expires.Format("2006-01-02 15:04:05"))
		if err == nil {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "没有权限",
			})
			c.Abort()
		}
	}
}
func AddCoin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("Id")
		if err != nil {
			fmt.Println("cookie读取错误")
		}
		todayTime := time.Now()
		today := strconv.Itoa(todayTime.Year()) + "year" + strconv.Itoa(todayTime.YearDay()) + "day"
		id, _ := strconv.Atoi(cookie)
		modles.AddCoin(id, today)
	}
}