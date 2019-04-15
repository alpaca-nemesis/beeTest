package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int
	Username string
	Password string
}

type result struct {
	Id			int
	title 		string
	class 		string
	content		string
	metadata	[]string
	date		time.Time
}

func init() {
	orm.RegisterModel(new(User), new(result))
}
