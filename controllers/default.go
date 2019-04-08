package controllers

import (
	"github.com/astaxie/beego"
)


/*************  INDEX  ******************/
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


/*************  SEARCH  ******************/
type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	c.TplName = "search.html"
}

func (c *SearchController) Post() {
	jsoninfo := c.GetString("content")
	c.Data["text"] = jsoninfo
	c.TplName = "search.html"
}


/*************  LOGIN  ******************/
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "search.html"
}

func (c *LoginController) Post() {
	c.TplName = "search.html"
}
