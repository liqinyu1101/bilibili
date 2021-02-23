package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func MysqlInit() *sql.DB {
	 //db, err := sql.Open("mysql", "root:dengjie123@tcp(127.0.0.1:3306)/bilibili?charset=utf8")
	db,err:=sql.Open("mysql","xiaoxu:xiaoxu@tcp(118.178.190.150:3306)/bilibili?charset=utf8")
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	DB = db
	return DB
}
