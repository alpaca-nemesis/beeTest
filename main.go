package main

import (
	_ "beeTest/models"
	_ "beeTest/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}


