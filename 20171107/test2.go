package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

/*
Go实现的支持PostgreSQL的驱动也很多，因为国外很多人在开发中使用了这个数据库。

https://github.com/lib/pq 支持database/sql驱动，纯Go写的
https://github.com/jbarham/gopgsqldriver 支持database/sql驱动，纯Go写的
https://github.com/lxn/go-pgsql 支持database/sql驱动，纯Go写的
*/
func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=test sslmode=disable")
	checkErr(err)

	//插入数据,PostgreSQL是通过$1,$2这种方式来指定要传递的参数
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("hello", "研发部门", "2017-12-09")
	checkErr(err)

	//pg不支持这个函数，因为他没有类似MySQL的自增ID
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "hello", "研发部门", "2017-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("最后插入id =", lastInsertId)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("test", lastInsertId)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department，created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(lastInsertId)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
