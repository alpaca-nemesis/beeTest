package controllers

import (
	_ "beeTest/models"
	"github.com/astaxie/beego"
)

/*************  INDEX  ******************/
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "localhost:8080/index"
	c.Data["Email"] = "mdzhangkailai@126.com"
	c.TplName = "index.tpl"
}

/*************  INDEX  ******************/
type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	islogin := c.Ctx.GetCookie("islogin")
	if islogin == "0" {
		c.Data["isLogin"] = 0
	} else {
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")
	}
	c.TplName = "index.html"
}

/*************  SEARCH  ******************/
type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	islogin := c.Ctx.GetCookie("islogin")
	if islogin == "0" {
		c.Data["isLogin"] = 0
		c.Redirect("/index", 302)
	}
	c.Data["isLogin"] = 1
	c.Data["username"] = c.GetSession("username")

	c.TplName = "result.html"
}


/*************  SEARCH  ******************/
type InterestingController struct {
	beego.Controller
}

func (c *InterestingController) Get() {
	c.TplName = "aaa.html"
}