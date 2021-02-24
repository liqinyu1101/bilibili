package modles

import (
	"bilibili/dao"
	"fmt"
)

type Comment struct {
	Username string `json:"username" form:"username"`
	Content string	`json:"content"  form:"content"`
	ReleaseTime string	`json:"releaseTime"`
	Bv string	`json:"bv" form:"bv"`
}
func PostComment(comment1 Comment)bool{
	stmt,err:=dao.DB.Prepare("insert into comment(username, content, releaseTime, bv) VALUES (?,?,?,?)")
	if err!=nil{
		fmt.Println("prepare failed",err)
		return false
	}
	_,err=stmt.Exec(comment1.Username,comment1.Content,comment1.ReleaseTime,comment1.Bv)
	if err!=nil{
		fmt.Println("exec failed",err)
		return false
	}
	return true
}
func ViesComment(bv string)[]Comment{
	stmt,err:=dao.DB.Query("select username, content, releaseTime, bv from comment where bv=? ",bv)
	if err!=nil{
		fmt.Println("query failed",err)
		return nil
	}
	var comments []Comment
	defer stmt.Close()
	for stmt.Next(){
		var comment Comment
		err=stmt.Scan(&comment.Username,&comment.Content,&comment.ReleaseTime,&comment.Bv)
		if err!=nil{
			fmt.Println("scam failed",err)
			return nil
		}
		comments= append(comments, comment)
	}
	return comments
}