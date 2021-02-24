package modles

import (
	"bilibili/dao"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Id       int    `form:"id"`
	Username string `form:"username"`
	Phone   string `form:"phone"`
	Coin    int    `form:"coin"`
	Follows int    `form:"follows"`
	Fans    int    `form:"fans"`
}
type Status struct {
	Coin	int
	ThumbsUp int
	Collect	int
}
func Register_(username string, password string, phone string) int {
	stmt, err := dao.DB.Prepare("insert into user(username,password,phone) values (?,?,?);")
	fmt.Println(username, password, phone)
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		return 0
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password, phone)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return 0
	}
	st, er := dao.DB.Query("select id from user where username=?;", username)
	if er != nil {
		return 0
		fmt.Printf("query failed:%v", er)
	}
	defer st.Close()
	var id int
	for st.Next() {
		er = st.Scan(&id)
		if er != nil {
			fmt.Printf("scan failed %v", er)
			return 0
		}
	}
	return id
	/*var u User
	for st.Next(){
		er=st.Scan(&u.Id,&u.Username,&u.Password,&u.Phone,&u.Qq,&u.Coin,&u.Follows,&u.Fans)
		if er!=nil{
			fmt.Printf("scan failed %v",err)
			return 0
		}
	}*/
}
func Login_(username string, phone string, password string) string {
	if strings.EqualFold(phone, "") {
		//用户名登录
		st, err := dao.DB.Query("select username,password,id from user where username=?;", username)
		if err != nil {
			return ""
			fmt.Printf("query failed:%v", err)
		}
		defer st.Close()
		var username, password1 string
		var id int
		for st.Next() {
			err = st.Scan(&username, &password1, &id)
			if err != nil {
				fmt.Printf("scan failed %v", err)
				return ""
			}
		}
		if strings.EqualFold(password1, password) {
			return strconv.Itoa(id)
		} else {
			return ""
		}

	} else {
		//手机号登录
		st, err := dao.DB.Query("select username,password,id from user where phone=?;", phone)
		if err != nil {
			return ""
			fmt.Printf("query failed:%v", err)
		}
		defer st.Close()
		var username, password1 string
		var id int
		for st.Next() {
			err = st.Scan(&username, &password1, &id)
			if err != nil {
				fmt.Printf("scan failed %v", err)
				return ""
			}
		}
		if strings.EqualFold(password, password1) {
			return strconv.Itoa(id)
		} else {
			return ""
		}

	}
}
func UserMessage(id string) User {
	u := User{}
	st, err := dao.DB.Query("select id,username,phone,coin,follows,fans from user where id=?;", id)
	if err != nil {
		return u
		fmt.Printf("query failed:%v", err)
	}
	defer st.Close()
	for st.Next() {
		err = st.Scan(&u.Id, &u.Username, &u.Phone, &u.Coin, &u.Follows, &u.Fans)
		if err != nil {
			fmt.Printf("scan failed %v", err)
			return u
		}
	}
	return u
}
func AddCoin(id int, today string) {
	st, err := dao.DB.Query("select coinDay from user where id=?;", id)
	if err != nil {
		fmt.Printf("query failed %v\n", err)
	}
	defer st.Close()
	var coinDay string
	for st.Next() {
		err = st.Scan(&coinDay)
		if err != nil {
			fmt.Printf("查询上次登录错误%v", err)
		}
	}
	if !strings.EqualFold(coinDay, today) {
		fmt.Println(coinDay, today)
		stmt, er := dao.DB.Prepare("update user set coinDay=? where id =?")
		if er != nil {
			fmt.Printf("硬币修改准备失败%v\n", er)
		}
		_, er = stmt.Exec(today, id)
		if er != nil {
			fmt.Printf("硬币修改失败%v\n", er)
		}
		st, err := dao.DB.Query("select coin from user where id=?", id)
		if err != nil {
			fmt.Println("query failed:", err)
		}
		defer st.Close()
		var coin int
		for st.Next() {
			err = st.Scan(&coin)
			if err != nil {
				fmt.Println("scan failed:", err)
			}
		}
		coin++
		stmt, er = dao.DB.Prepare("update user set coin=? where id=?")
		if er != nil {
			fmt.Println("coin update failed", er)
		}
		_, er = stmt.Exec(coin, id)
		if er != nil {
			fmt.Println("coin update failed", er)
		}
	}

}

//1取消关注成功，0未知错误，-1关注成功
func Follow(fan, up string) int {
	st, er := dao.DB.Query("select id from follow where fan=? and  up=?", fan, up)
	if er != nil {
		fmt.Println("follow query failed", er)
		return 0
	}
	flag := -1
	defer st.Close()
	for st.Next() {
		er := st.Scan(&flag)
		if er != nil {
			fmt.Println("follow failed", er)
			return 0

		}
	}
	if flag != -1 {
		flag=1
		stmt,err:=dao.DB.Prepare("delete from follow where fan=? and up=?")
		if err!=nil{
			fmt.Println("prepare failed", err)
			return 0
		}
		_,err=stmt.Exec(fan,up)
		if err!=nil{
			fmt.Println("exec failed", err)
		}
	} else {
		stmt, err := dao.DB.Prepare("insert into follow(fan, up) VALUES (?,?)")
		if err != nil {
			fmt.Println("prepare failed", err)
			return 0
		}
		_, err = stmt.Exec(fan, up)
		if err != nil {
			fmt.Println("exec failed", err)
			return 0
		}
	}
	return flag
}
func SearchUser(keyword string)[]User{
	stmt,err:=dao.DB.Query("select id,username,phone,coin,follows,fans from user where user.id like ? or user.username like ?","%"+keyword+"%","%"+keyword+"%")
	var users []User
	if err!=nil{
		fmt.Println("query failed",err)
		return nil
	}

	for stmt.Next(){
	var u User
		err=stmt.Scan(&u.Id,&u.Username,&u.Phone,&u.Coin,&u.Follows,&u.Fans)
		if err!=nil{
			fmt.Printf("scan failed %v",err)
			return nil
		}
		users=append(users,u)
	}
	return users
}
func FollowChange(follow int,id string)bool{
	stmt,err:=dao.DB.Prepare("update user set follows = ? where id = ?")
	if err!=nil{
		fmt.Println("prepare failed",err)
	}
	_,err=stmt.Exec(follow,id)
	if err!=nil{
		fmt.Println("exec failed",err)
	}
	return true
}
func FanChange(fans int,id string) bool{
	stmt,err:=dao.DB.Prepare("update user set fans = ? where id = ?")
	if err!=nil{
		fmt.Println("prepare failed",err)
	}
	_,err=stmt.Exec(fans,id)
	if err!=nil{
		fmt.Println("exec failed",err)
	}
	return true
}
func RegisterUsername(username string)bool {
	stmt,err:=dao.DB.Query("select id from user where username = ? ",username)
	if err!=nil{
		fmt.Println("prepare failed",err)
		return false
	}
	var flag int
	defer stmt.Close()
	for stmt.Next(){
		err=stmt.Scan(&flag)
		if err!=nil{
			fmt.Println("scan failed",err)
			return false
		}
	}
	if flag==0{
		return true
	}
	return false
}
func RegisterPhone(phone string)bool {
	stmt,err:=dao.DB.Query("select id from user where phone = ? ",phone)
	if err!=nil{
		fmt.Println("prepare failed",err)
		return false
	}
	var flag int
	defer stmt.Close()
	for stmt.Next(){
		err=stmt.Scan(&flag)
		if err!=nil{
			fmt.Println("scan failed",err)
			return false
		}
	}
	if flag==0{
		return true
	}
	return false
}