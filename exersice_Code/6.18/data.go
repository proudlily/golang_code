package main

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"fmt"
	//"time"
)

func main() {
	db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}
