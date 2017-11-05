package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //只加载，不使用,安装: go get github.com/go-sql-driver/mysql
)

//Go没有内置的驱动支持任何的数据库，但是Go定义了database/sql接口，用户可以基于驱动接口开发相应数据库的驱动
//see :http://go-database-sql.org/
//database/sql在database/sql/driver提供的接口基础上定义了一些更高阶的方法
//用以简化数据库操作,同时内部还建议性地实现一个conn pool。
/*
Go中支持MySQL的驱动目前比较多，有如下几种，有些是支持database/sql标准，而有些是采用了自己的实现接口,
常用的有如下几种:
https://github.com/go-sql-driver/mysql 支持database/sql，全部采用go写。
https://github.com/ziutek/mymysql 支持database/sql，也支持自定义的接口，全部采用go写。
https://github.com/Philio/GoMySQL 不支持database/sql，自定义接口，全部采用go写。
推荐使用第一个：
1.这个驱动比较新，维护的比较好
2.完全支持database/sql接口
3.支持keepalive，保持长连接,虽然星星fork的mymysql也支持keepalive，但不是线程安全的，这个从底层就支持了keepalive。
*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
db.Query()函数用来直接执行Sql返回Rows结果。
stmt.Exec()函数用来执行stmt准备好的SQL语句
*/
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("insert userinfo set username=? ,departname=?,created=? ")
	//stmt, err := db.Prepare("insert into userinfo(username,departname,created)values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("tom", "研发部", "2017-10-11")
	checkErr(err)
	id, err := res.LastInsertId()
	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? ,departname=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("GOGOGO", "++++++", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var ids []int
	for rows.Next() {
		var id int
		var username string
		var departname string
		var created string
		err = rows.Scan(&id, &username, &departname, &created)
		checkErr(err)
		fmt.Println(id, username, departname, created)
		ids = append(ids, id)
	}
	fmt.Println(ids)

	//删除数据

	if len(ids) > 0 {
		for _, id := range ids {
			stmt, err = db.Prepare("delete from userinfo where uid = ? ")
			checkErr(err)
			res, err = stmt.Exec(id)
			checkErr(err)
			affect, err = res.RowsAffected()
			checkErr(err)
			fmt.Println("delete row count : ", affect)
		}
	}
	db.Close()
}
