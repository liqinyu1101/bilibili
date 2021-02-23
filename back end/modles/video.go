package modles

import (
	"bilibili/dao"
	"fmt"
)

type Video struct {
	Bv        int
	Title     string
	Commit    string
	Coin      int
	ThumbsUp      int
	Author    string
	ReleaseTime      string
	Collect   int
	Times     int //观看次数
	Tag       string
}

func MaxBv()int{
	stmt,err:=dao.DB.Query("select max(bv) from  video")
	if err!=nil{
		fmt.Println("query failed",err)
	}
	var maxBv int
	defer stmt.Close()
	for stmt.Next(){
		err=stmt.Scan(&maxBv)
		if err!=nil{
			fmt.Println("scan failed",err)
		}
	}
	return maxBv
}
func UploadVideo(title,  author,commit,ReleaseTime,tag string)bool{
	stmt,err:=dao.DB.Prepare("insert into video(title, author,commit,ReleaseTime,tag) values (?,?,?,?,?)")
	if err!=nil{
		fmt.Println("prepare failed",err)
		return false
	}
	_,err=stmt.Exec(title,author,commit,ReleaseTime,tag)
	if err!=nil{
		fmt.Println("insert failed",err)
		return false
	}
	return true
}
func RecommendVideoTag(number int,tag string,)[]Video{
	stmt,err:=dao.DB.Query("select * from video where tag = ? order by times desc limit ?",tag,number)
	if err!=nil{
		fmt.Println("query failed",err)
		return nil
	}
	defer stmt.Close()
	var videos []Video
	for stmt.Next(){
		var video Video
		err=stmt.Scan(&video.Bv,&video.Title,&video.Author,&video.ReleaseTime,&video.Collect,&video.ThumbsUp,&video.Coin,&video.Times,&video.Commit,&video.Tag)
		if err!=nil{
			fmt.Println("scan failed",err)
			return nil
		}
		videos=append(videos,video)
	}
	return videos
}
func RecommendVideoTimes(number int)[]Video{
	stmt,err:=dao.DB.Query("select * from video  order by times desc limit ?",number)
	if err!=nil{
		fmt.Println("query failed",err)
		return nil
	}
	defer stmt.Close()
	var videos []Video
	for stmt.Next(){
		var video Video
		err=stmt.Scan(&video.Bv,&video.Title,&video.Author,&video.ReleaseTime,&video.Collect,&video.ThumbsUp,&video.Coin,&video.Times,&video.Commit,&video.Tag)
		if err!=nil{
			fmt.Println("scan failed",err)
			return nil
		}
		videos=append(videos,video)
	}
	return videos
}
func VideoByBv(bv string)Video{
	var video Video
	stmt, err := dao.DB.Query("select * from video where bv=?",bv)
	if err!=nil{
		fmt.Println("query failed",err)
	}
	defer stmt.Close()
	for stmt.Next(){
		err=stmt.Scan(&video.Bv, &video.Title, &video.Author, &video.ReleaseTime, &video.Collect, &video.ThumbsUp, &video.Coin, &video.Times, &video.Commit, &video.Tag)
		if err!=nil{
			fmt.Println("scan failed",err)
		}
	}
	return video
}
//1成功，0未知错误，-1已经投过币，-2硬币不足
func InsertCoin(id,bv string) int {
	flag:=InsertCoinOrNot(id,bv)
	if flag!=-1{
		return -1
	}else {
		user_coin:=UserMessage(id).Coin
			if user_coin==0{
				return -2
			}else {
				bv_coin:=VideoByBv(bv).Coin
				st,er:=dao.DB.Prepare("update user set coin=? where id=?")
				if er!=nil{
					fmt.Println("prepare wrong ",er)
					return 0
				}
				user_coin--
				_, er =st.Exec(user_coin,id)
				if er!=nil{
					fmt.Println("update wrong",er)
					return 0
				}
				st,er=dao.DB.Prepare("insert into coin(bv, id) VALUES (?,?)")
				if er!=nil{
					fmt.Println("prepare wrong ",er)
					return 0
				}
				_, er =st.Exec(bv,id)
				if er!=nil{
					fmt.Println("insert wrong",er)
					return 0
				}
				st,er=dao.DB.Prepare("update video set coin=? where bv=?")
				if er!=nil{
					fmt.Println("prepare wrong ",er)
					return 0
				}
				bv_coin++
				_, er =st.Exec(bv_coin,bv)
				if er!=nil{
					fmt.Println("update wrong",er)
					return 0
				}
				return 1
			}
		}

}


func SearchVideo(keyword string)[]Video {
	stmt, err := dao.DB.Query("select * from video where video.title like ? or video.author like ? or video.commit like ?", "%" + keyword + "%","%" + keyword + "%", "%" + keyword + "%")
	if err != nil {
		fmt.Println("query failed", err)
	}
	defer stmt.Close()
	var videos []Video
	for stmt.Next() {
		var video Video
		err = stmt.Scan(&video.Bv, &video.Title, &video.Author, &video.ReleaseTime, &video.Collect, &video.ThumbsUp, &video.Coin, &video.Times, &video.Commit, &video.Tag)
		if err != nil {
			fmt.Println("scan failed", err)
			return nil
		}
		videos = append(videos, video)
	}
	return videos
}

//-1未点赞，0未知错误，正数已点赞
func ThumbsUpOrNot(id,bv string) int{
	stmt,err:=dao.DB.Query("select flag from ThumbsUp where bv=? and id=?",bv,id)
	if err!=nil{
		fmt.Println("query failed",err)
		return 0
	}
	flag:=-1
	defer stmt.Close()
	for stmt.Next(){
		err=stmt.Scan(&flag)
		if err!=nil{
			fmt.Println("scan failed",err)
		}
	}
	return flag
}
//1 已经点赞 0未知错误 -1 取消点赞
func ThumbsUp(id,bv string)int{
	flag:=ThumbsUpOrNot(id,bv)
	likes:=VideoByBv(bv).ThumbsUp
	if flag!=-1{
		st,er:=dao.DB.Prepare("delete from ThumbsUp where flag=?")
		if er!=nil{
			fmt.Println("prepare failed ",er)
			return 0
		}
		_,er=st.Exec(flag)
		if er!=nil{
			fmt.Println("delete failed ",er)
			return 0
		}
		likes--
		st,er=dao.DB.Prepare("update video set likes=? where bv=?")
		if er!=nil{
			fmt.Println("prepare failed ",er)
			return 0
		}
		_,er=st.Exec(likes,bv)
		if er!=nil{
			fmt.Println("delete failed ",er)
			return 0
		}
		flag=-1
	}else {

		flag=1
		st,er:=dao.DB.Prepare("insert into ThumbsUp(bv ,id) values (?,?)")
		if er!=nil{
			fmt.Println("prepare failed ",er)
			return 0
		}
		_,er=st.Exec(bv,id)
		if er!=nil{
			fmt.Println("insert failed ",er)
			return 0
		}
		likes++
		st,er=dao.DB.Prepare("update video set likes=? where bv=?")
		if er!=nil{
			fmt.Println("prepare failed ",er)
			return 0
		}
		_,er=st.Exec(likes,bv)
		if er!=nil{
			fmt.Println("delete failed ",er)
			return 0
		}
		return flag
	}
	return flag
}
//-1未关注，0未知错误，正数已关注
func CollectOrNot(id ,bv string)int{
	stmt,err:=dao.DB.Query("select flag from Collect where bv=? and id=?",bv,id)
	if err!=nil{
		fmt.Println("query failed",err)
		return 0
	}
	flag:=-1
	defer stmt.Close()
	for stmt.Next(){
		err=stmt.Scan(&flag)
		if err!=nil{
			fmt.Println("scan failed",err)
		}
	}
	return flag
}
//1 收藏成功 0未知错误 -1 取消收藏
func Collect(id,bv string)int{
	flag:=CollectOrNot(id,bv)
	collect:=VideoByBv(bv).Collect
	if flag!=-1{
		st,er:=dao.DB.Prepare("delete from collect where flag=?")
		if er!=nil{
			fmt.Println("prepare failed ",er)
			return 0
		}
		_,er=st.Exec(flag)
		if er!=nil{
			fmt.Println("delete failed ",er)
			return 0
		}
		collect--
		st,er=dao.DB.Prepare("update video set collect=? where bv=?")
		if er!=nil{
			fmt.Println("prepare failed ",er)
			return 0
		}
		_,er=st.Exec(collect,bv)
		if er!=nil{
			fmt.Println("delete failed ",er)
			return 0
		}
		flag=-1
	}else {

		flag=1
		st,er:=dao.DB.Prepare("insert into collect(bv ,id) values (?,?)")
		if er!=nil{
			fmt.Println("prepare failed ",er)
			return 0
		}
		_,er=st.Exec(bv,id)
		if er!=nil{
			fmt.Println("insert failed ",er)
			return 0
		}
		collect++
		st,er=dao.DB.Prepare("update video set collect=? where bv=?")
		if er!=nil{
			fmt.Println("prepare failed ",er)
			return 0
		}
		_,er=st.Exec(collect,bv)
		if er!=nil{
			fmt.Println("delete failed ",er)
			return 0
		}
		return flag
	}
	return flag
}
func InsertCoinOrNot(id,bv string) int{
	stmt,err:=dao.DB.Query("select flag from coin where id=? and bv=?",id,bv)
	if err!=nil{
		fmt.Println("query failed",err)
		return 0
	}
	flag:=-1
	defer stmt.Close()
	for stmt.Next(){
		err=stmt.Scan(&flag)
		if err!=nil{
			fmt.Println("scan failed",err)
			return 0
		}
	}
	return flag
}
func ViewPostVideo(author string)[]Video{
	stmt,err:=dao.DB.Query("select * from video where author = ?",author)
	if err!=nil{
		fmt.Println("query failed",err)
		return nil
	}
	defer stmt.Close()
	var videos []Video
	for stmt.Next(){
		var video Video
		err=stmt.Scan(&video.Bv,&video.Title,&video.Author,&video.ReleaseTime,&video.Collect,&video.ThumbsUp,&video.Coin,&video.Times,&video.Commit,&video.Tag)
		if err!=nil{
			fmt.Println("scan failed",err)
			return nil
		}
		videos=append(videos,video)
	}
	fmt.Println(videos)
	return videos
}
func VideoTimesAdd(bv string,times int) bool {
	stmt,err:=dao.DB.Prepare("update video set times = ? where bv = ? ")
	if err!=nil{
		fmt.Println("Prepare failed",err)
	}
	_,err=stmt.Exec(times,bv)
	if err!=nil{
		fmt.Println("exec failed",err)
	}
	return true
}
func DeleteVideo(bv string) bool{
	stmt,err:=dao.DB.Prepare("delete from video where bv =?")
	if err!=nil{
		fmt.Println("prepare failed",err)
		return false
	}
	_,err=stmt.Exec(bv)
	if err!=nil{
		fmt.Println("exec failed",err)
		return false
	}
	return true
}