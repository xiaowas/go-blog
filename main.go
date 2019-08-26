package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "leechan.inline/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:2014gaokao@tcp(127.0.0.1:3306)/blog?charset=utf8&loc=Asia%2FShanghai")
}

func main() {
	beego.Run()
}
