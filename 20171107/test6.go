package main

import (
	"fmt"
	orm "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
)

//beego orm 采用了Go style方式对数据库进行操作，实现了struct到数据表记录的映射。
//目前开源的Go ORM框架中实现比较完整的一个库，而且运行效率相当不错，功能也基本能满足需求。
//beego orm是支持database/sql标准接口的ORM库，所以理论上来说，只要数据库驱动支持database/sql
//接口就可以无缝的接入beego orm。

//支持数据库:Mysql,PostgreSQL,SQLite
/*
驱动包括下面几个：
Mysql: github/go-mysql-driver/mysql
PostgreSQL: github.com/lib/pq
SQLite: github.com/mattn/go-sqlite3
Mysql: github.com/ziutek/mymysql/godrv

安装：go get github.com/astaxie/beego
*/
//Model
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	//注册定义的model
	orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))
	// 创建table
	orm.RunSyncdb("default", false, true)
}

/*
PostgreSQL 配置:
// _ "github.com/lib/pq"
// 注册驱动
orm.RegisterDriver("postgres", orm.DR_Postgres)
// 设置默认数据库
//PostgresQL用户：postgres ，密码：zxxx ， 数据库名称：test ， 数据库别名：default
orm.RegisterDataBase("default", "postgres", "user=postgres password=zxxx dbname=test host=127.0.0.1 port=5432 sslmode=disable")
MySQL 配置:
//导入驱动
//_ "github.com/go-sql-driver/mysql"
//注册驱动
orm.RegisterDriver("mysql", orm.DR_MySQL)
// 设置默认数据库
//mysql用户：root ，密码：zxxx ， 数据库名称：test ， 数据库别名：default
 orm.RegisterDataBase("default", "mysql", "root:zxxx@/test?charset=utf8")
Sqlite 配置:
//导入驱动
//_ "github.com/mattn/go-sqlite3"
//注册驱动
orm.RegisterDriver("sqlite", orm.DR_Sqlite)
// 设置默认数据库
//数据库存放位置：./datas/test.db ， 数据库别名：default
orm.RegisterDataBase("default", "sqlite3", "./datas/test.db")
*/
func main() {
	o := orm.NewOrm()
	user := User{Name: "hello"}
	id, err := o.Insert(&user) //InsertMulti 插入多条
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("id=", id)
	// 插入表
	user.Name = "Jack"
	num, err := o.Update(&user) // o.Update(&user, "Field1", "Field2", ...) 指定更新字段
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	fmt.Println(u)
	// 删除
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

/*
根据数据库的别名，设置数据库的最大空闲连接
orm.SetMaxIdleConns("default", 30)
根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)
orm.SetMaxOpenConns("default", 30)
目前beego orm支持打印调试，你可以通过如下的代码实现调试
orm.Debug = true

err:
if err == orm.ErrNoRows {
	fmt.Println("查询不到")
} else if err == orm.ErrMissPK {
	fmt.Println("找不到主键")
}

Where:用来设置条件，支持多个参数，第一个参数如果为整数，相当于调用了Where("主键=?",值)。
var user User
qs := o.QueryTable(user) // 返回 QuerySeter
qs.Filter("id", 1) // WHERE id = 1
qs.Filter("profile__age", 18) // WHERE profile.age = 18

qs.Filter("profile__age__in", 18, 20)  //WHERE profile.age IN (18, 20)
qs.Filter("profile__age__in", 18, 20).Exclude("profile__lt", 1000) // WHERE profile.age IN (18, 20) AND NOT profile_id < 1000
qs.Filter("profile__age__gt", 17) //WHERE profile.age > 17

qs.Limit(10, 20)  //LIMIT 10 OFFSET 20 注意跟SQL反过来的

qs.OrderBy("id", "-profile__age") // ORDER BY id ASC, profile.age DESC
qs.OrderBy("-profile__age", "profile") // ORDER BY profile.age DESC, profile_id ASC

GroupBy:用来指定进行groupby的字段
Having:用来指定having执行的时候的条件

关联查询:
type Post struct {
	Id    int    `orm:"auto"`
	Title string `orm:"size(100)"`
	User  *User  `orm:"rel(fk)"`
}

var posts []*Post
qs := o.QueryTable("post")
num, err := qs.Filter("User__Name", "slene").All(&posts)

使用原生sql
var r RawSeter
r = o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")

//查询
func (m *User) Query(name string) []User {
var o orm.Ormer
var rs orm.RawSeter
o = orm.NewOrm()
rs = o.Raw("SELECT * FROM user "+
	"WHERE name=? AND uid>10 "+
	"ORDER BY uid DESC "+
	"LIMIT 100", name)
var user []User
num, err := rs.QueryRows(&user)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(num)
	return user
}
}
*/
