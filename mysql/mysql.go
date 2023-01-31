package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func initMySQL() (err error) {
	var dbuser string
	fmt.Println("用户：")
	fmt.Scanln(&dbuser)
	var dbpasswd string
	fmt.Println("密码：")
	fmt.Scanln(&dbpasswd)
	var dbhost string
	fmt.Println("ip：")
	fmt.Scanln(&dbhost)
	var dbport string
	fmt.Println("端口：")
	fmt.Scanln(&dbport)
	var dbname string
	fmt.Println("数据库名称：")
	fmt.Scanln(&dbname)
	dsn := dbuser + ":" + dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(200)                 //最大连接数
	db.SetMaxIdleConns(10)                  //连接池里最大空闲连接数。必须要比maxOpenConns小
	db.SetConnMaxLifetime(time.Second * 10) //最大存活保持时间
	db.SetConnMaxIdleTime(time.Second * 10) //最大空闲保持时间
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
	} else {
		fmt.Println("connect to db success")
	}
	defer db.Close()
}
