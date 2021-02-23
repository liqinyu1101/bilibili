package service

import (
	"bilibili/modles"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func Register_(ctx *gin.Context) int {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	res := modles.Register_(username, password, phone)
	if res != 0 {
		//记录登录状态
		cookie := &http.Cookie{
			Name:     "Id",
			Value:    strconv.Itoa(res),
			Path:     "/",
			HttpOnly: false,
			MaxAge:   2000,
		}
		http.SetCookie(ctx.Writer, cookie)
	}
	return res
}
func Login_(ctx *gin.Context) bool {
	username := ctx.PostForm("username")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")
	res := modles.Login_(username, phone, password)
	if strings.EqualFold(res, "") {
		return false
	} else {
		//记录登录状态
		cookie := &http.Cookie{
			Name:     "Id",
			Value:    res,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   2000,
		}
		http.SetCookie(ctx.Writer, cookie)
		return true
	}
}

func SelfMessage(ctx *gin.Context) modles.User {
	u := modles.User{}
	cookie, err := ctx.Request.Cookie("Id")
	if err != nil {
		fmt.Println("Haven't login")
	} else {
		id := cookie.Value
		u = modles.UserMessage(id)
	}
	return u
}
func UserMessage(ctx *gin.Context) modles.User {
	var u modles.User
	id:=ctx.PostForm("id")
	u = modles.UserMessage(id)
	return u
}
func UploadFace(ctx *gin.Context) bool {
	face, err := ctx.FormFile("face")
	if err != nil {
		fmt.Printf("formfile failed+%v", err)
		return false
	} else {
		//用用户id当作文件路径
		cookie, _ := ctx.Request.Cookie("Id")
		des := path.Join("./source/face", cookie.Value+".jpg")
		err = ctx.SaveUploadedFile(face, des)
		if err != nil {
			fmt.Printf("SaveUpLoadedFile failed:%v", err)
			return false
		} else {
			return true
		}
	}
}
//1取消关注成功，0未知错误，-1关注成功
func Follow(ctx *gin.Context)int{
	fanId,err:=ctx.Cookie("Id")
	if err!=nil{
		return 0
	}
	fan:=modles.UserMessage(fanId)
	upId:=ctx.PostForm("up")
	res:=modles.Follow(fanId,upId)
	if res==-1{
		//关注成功
		fanFollow:=fan.Follows
		fanFollow++
		if !modles.FollowChange(fanFollow,fanId){
			return 0
		}
		upFans:=modles.UserMessage(upId).Fans
		upFans++
		if !modles.FanChange(upFans,upId){
			return 0
		}
	}else if res ==1 {
		fanFollow:=fan.Follows
		fanFollow--
		if !modles.FollowChange(fanFollow,fanId){
			return 0
		}
		upFans:=modles.UserMessage(upId).Fans
		upFans--
		if !modles.FanChange(upFans,upId){
			return 0
		}
	}
	return res
}
func SearchUser(ctx *gin.Context)string{
	keyword:=ctx.Query("keyword")
	users:=modles.SearchUser(keyword)
	videosJson,err:=json.Marshal(users)
	if err!=nil{
		fmt.Println("json wrong",err)
	}
	return string(videosJson)
}
func VideoStatus(ctx *gin.Context)modles.Status{
	var status modles.Status
	bv:=ctx.PostForm("bv")
	id,err:=ctx.Cookie("Id")
	if err!=nil{
		fmt.Println("read cookie",err)
	}
	status.ThumbsUp=modles.ThumbsUpOrNot(id,bv)
	status.Collect=modles.CollectOrNot(id,bv)
	status.Coin=modles.InsertCoinOrNot(id,bv)
	return status
}
func RegisterUsername(ctx *gin.Context)bool {
	username:=ctx.PostForm("username")
	res:=modles.RegisterUsername(username)
	return res
}
func RegisterPhone(ctx *gin.Context)bool {
	phone:=ctx.PostForm("phone")
	res:=modles.RegisterPhone(phone)
	return res
}