package controllers

import (
	_ "beeTest/models"
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	"log"
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
	islogin := c.GetSession("islogin")
	if islogin == nil {
		c.Data["isLogin"] = 0
	} else {
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")
	}
	c.TplName = "index.html"
}

/*************  SEARCH  ******************/
type SearchController struct {
	client *elastic.Client
	beego.Controller
}



func (c *SearchController) Get() {
	islogin := c.GetSession("islogin")
	if islogin == nil {
		c.Data["isLogin"] = 0
		c.Redirect("/index", 302)
	} else {
		if c.client == nil {
			err := c.clientInit()
			if err != nil {
				log.Fatalln("client err: ", err)
			}
		}
		var content string
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")
		err := c.Ctx.Input.Bind(&content, "content")
		if err != nil {
			log.Fatalln("content nil: ", err)
		}
		var result string
		result = c.searchContent(content)
		c.Data["result"] = result

		c.Data["searchContent"] = content
		c.TplName = "result.html"
	}
}


/*************  AddContent  ******************/
type AddContentController struct {
	client *elastic.Client
	beego.Controller
}

func (c *AddContentController) Get() {
	islogin := c.GetSession("islogin")
	if islogin == nil {
		c.Data["isLogin"] = 0
		c.Redirect("/index", 302)
	} else {
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")
		c.TplName = "add.html"
	}
}

func (c *AddContentController) Post() {
	islogin := c.GetSession("islogin")
	if islogin == nil {
		c.Data["isLogin"] = 0
		c.Redirect("/index", 302)
	} else {
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")
		c.TplName = "add.html"
	}
}


/*************  Interesting  ******************/
type InterestingController struct {
	beego.Controller
}

func (c *InterestingController) Get() {
	c.TplName = "relax.html"
}


/*************  CXKBALL  ******************/
type CXKController struct {
	beego.Controller
}

func (c *CXKController) Get() {
	c.TplName = "cxk.html"
}

