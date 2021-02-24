package modles

import (
	"bilibili/dao"
	"fmt"
)

type Bullet struct {
	TimePoint string `json:"time" form:"time"`
	Content string	`json:"content" form:"content"`
	Bv string `json:"bv" form:"bv"`
}
func ViewBullet(bv string)[]Bullet {
	stmt,err:=dao.DB.Query("select timePoint, content ,bv from bullet where bv = ?",bv)
	if err!=nil{
		fmt.Println("query failed",err)
		return nil
	}
	defer stmt.Close()
	var  bullets []Bullet
	for stmt.Next(){
		var bullet Bullet
		err=stmt.Scan(&bullet.TimePoint,&bullet.Content,&bullet.Bv)
		bullets=append(bullets,bullet)
		if err!=nil{
			fmt.Println("scan failed",err)
		}
	}
	return bullets
}
func PostBullet(bullet Bullet)bool{
	stmt,err:=dao.DB.Prepare("insert into bullet( timePoint, content, bv) VALUES (?,?,?)")
	if err!=nil{
		fmt.Println("prepare failed",err)
		return false
	}
	_,err=stmt.Exec(bullet.TimePoint,bullet.Content,bullet.Bv)
	if err!=nil{
		fmt.Println("exec failed",err)
		return false
	}
	return true
}