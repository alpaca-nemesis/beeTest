package main

import (
	_ "beeTest/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
