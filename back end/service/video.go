package service

import (
	"bilibili/modles"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)


func UploadVideo(ctx *gin.Context) bool {
	title:=ctx.PostForm("title")
	commit:=ctx.PostForm("commit")
	tag:=ctx.PostForm("tag")
	authorId,err:=ctx.Cookie("Id")
	if err!=nil{
		fmt.Println("cookie wrong",err)
	}
	author:=modles.UserMessage(authorId).Username
	ReleaseTime:=time.Now().Format("2006-01-02 15:04:05")
	res:=modles.UploadVideo(title,author,commit,ReleaseTime,tag)
	bv:=strconv.Itoa(modles.MaxBv())
	if res{
		cover,err:=ctx.FormFile("cover")
		video,err:=ctx.FormFile("video")
		if err!=nil{
			fmt.Println("ger cover failed",err)
		}
		coverDst := path.Join("./source/cover",bv+".jpg")
		videoDst := path.Join("./source/video",bv+".mp4")
		err=ctx.SaveUploadedFile(cover,coverDst)
		if err!=nil{
			fmt.Println("cover save failed")
			return false
		}
		err=ctx.SaveUploadedFile(video,videoDst)
		if err!=nil{
			fmt.Println("video save failed")
			return false
		}
		return res
	}else {
		return false
	}
}
func RecommendVideo(ctx *gin.Context)string{
	tag:=ctx.PostForm("tag")
	number_:=ctx.PostForm("number")
	number,_:=strconv.Atoi(number_)
	var videos []modles.Video
	if strings.EqualFold(tag,""){
		videos=modles.RecommendVideoTimes(number)
	}else {
		videos=modles.RecommendVideoTag(number,tag)
	}
	videoString,err:=disableEscapeHtml(videos)
	if err!=nil{
		fmt.Println("转义问题",err)
	}
	return videoString
}
func disableEscapeHtml(data interface{}) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(&data); err != nil {
		return "", err
	}
	return bf.String(), nil
}
func InsertCoin(ctx *gin.Context)int{
	bv:=ctx.PostForm("bv")
	id,err:=ctx.Cookie("Id")
	if err!=nil{
		fmt.Println("read cookie",err)
	}
	res:=modles.InsertCoin(id,bv)
	return res
}
func ThumbsUp(ctx *gin.Context)int{
	bv:=ctx.PostForm("bv")
	id,err:=ctx.Cookie("Id")
	if err!=nil{
		fmt.Println("read cookie",err)
	}
	res:=modles.ThumbsUp(id,bv)
	return res
}
func SearchVideo(ctx *gin.Context)string{
	keyword:=ctx.Query("keyword")
	videos:=modles.SearchVideo(keyword)
	videosJson,err:=json.Marshal(videos)
	if err!=nil{
		fmt.Println("json wrong",err)
	}
	return string(videosJson)
}
//1 已经收藏 0未知错误 -1 取消收藏
func Collect(ctx *gin.Context)int{
	bv:=ctx.PostForm("bv")
	id,err:=ctx.Cookie("Id")
	if err!=nil{
		fmt.Println("read cookie",err)
	}
	res:=modles.Collect(id,bv)
	return res
}
func Links(ctx *gin.Context)bool{
	bv:=ctx.PostForm("bv")
	id,err:=ctx.Cookie("Id")
	if err!=nil{
		fmt.Println("read cookie",err)
	}
	flagCollect:=modles.CollectOrNot(id,bv)
	if flagCollect==-1{
		flagCollect=modles.Collect(id,bv)
	}
	if flagCollect<1{
		return false
	}
	flagCoin:=modles.InsertCoin(id,bv)
	if flagCoin==0 || flagCoin==-2{
		return false
	}
	flagThumbsUp:=modles.ThumbsUpOrNot(id,bv)
	if flagThumbsUp==-1{
		flagThumbsUp=modles.ThumbsUp(id,bv)
	}
	if flagThumbsUp<1 {
		return false
	}
	return true
}
func ViewPostVideo(ctx *gin.Context) string {
	authorId:=ctx.PostForm("id")
	author:=modles.UserMessage(authorId).Username
	videos:=modles.ViewPostVideo(author)
	videosJson,er:=disableEscapeHtml(videos)
	if er!=nil{
		fmt.Println("marshal failed",er)
	}
	return videosJson
}
func VideoTimesAdd(ctx *gin.Context)bool{
	bv:=ctx.PostForm("bv")
	times:=modles.VideoByBv(bv).Times
	times++
	res:=modles.VideoTimesAdd(bv,times)
	return res
}
func DeleteVideo(ctx *gin.Context)bool{
	bv:=ctx.PostForm("bv")
	authorReal:=modles.VideoByBv(bv).Author
	authorId,err:=ctx.Cookie("Id")
	if err!=nil{
		fmt.Println("cookie wrong",err)
	}
	author:=modles.UserMessage(authorId).Username
	if strings.EqualFold(author,authorReal){
		res:=modles.DeleteVideo(bv)
		coverDst := path.Join("./source/cover",bv+".jpg")
		videoDst := path.Join("./source/video",bv+".mp4")
		err:=os.Remove(coverDst)
		if err!=nil{
			fmt.Println("remove failed",err)
		}
		if err!=nil &&!res{
			return false
		}
		os.Remove(videoDst)
		return true
	}else {
		return false
	}
}
func ThumbsUpAndTimes(ctx *gin.Context)string{
	authorId:=ctx.PostForm("id")
	author:=modles.UserMessage(authorId).Username
	videos:=modles.ViewPostVideo(author)
	var times,Thumbs int
	for _,v:= range videos{
		times+=v.Times
		Thumbs+=v.ThumbsUp
	}
	type TAndT struct {
		Times int
		ThumbsUp int
	}
	a:= TAndT{
		ThumbsUp: Thumbs,
		Times: times,
	}
	message,_:=disableEscapeHtml(a)
	return message
}
