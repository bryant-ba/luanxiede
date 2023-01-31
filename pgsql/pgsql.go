package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=1234qwer dbname=postgres sslmode=disable")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("1234qwer", "ops", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
