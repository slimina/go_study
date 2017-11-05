package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

/*
Go支持sqlite的驱动也比较多，但是好多都是不支持database/sql接口的
https://github.com/mattn/go-sqlite3 支持database/sql接口，基于cgo
https://github.com/feyeleanor/gosqlite3 不支持database/sql接口，基于cgo写的
https://github.com/phf/go-sqlite3 不支持database/sql接口，基于cgo写的
目前支持database/sql的SQLite数据库驱动只有第一个，我目前也是采用它来开发项目的。
采用标准接口有利于以后出现更好的驱动的时候做迁移。

sqlitle3是个cgo库，需要gcd编译c代码,gcc
window 下载安装https://sourceforge.net/projects/mingw-w64/
或http://tdm-gcc.tdragon.net/download
并添加环境变量
*/
func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	checkErr(err)
	//添加数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("TOM", "研发部门", "2017-11-06")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更细数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err := stmt.Exec("Jack", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	//查询
	rows, err := db.Query("select * from userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department，created)
	}
	//删除
	stmt,err = db.Prepare("delect from userinfo where uid=?")
	checkErr(err)
	res,err = stmt.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
