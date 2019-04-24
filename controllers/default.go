package controllers

import (
	_ "beeTest/models"
	"github.com/astaxie/beego"
	"log"
	"net/rpc"
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
	rpcClient *rpc.Client
	beego.Controller
}

func (c *SearchController) Get() {
	islogin := c.GetSession("islogin")
	if c.rpcClient == nil{
		c.rpcClient = rpcInit()
	}
	if islogin == nil {
		c.Data["isLogin"] = 0
		c.Redirect("/index", 302)
	} else {
		var content string
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")
		err := c.Ctx.Input.Bind(&content, "content")
		if err != nil {
			log.Fatalln("content nil: ", err)
		}
		c.Data["searchContent"] = content
		c.Data["content"] = search(content, c.rpcClient)
		c.TplName = "result.html"
	}
}


/*************  Interesting  ******************/
type InterestingController struct {
	beego.Controller
}

func (c *InterestingController) Get() {
	c.TplName = "relax.html"
}

/*************  AddContent  ******************/
type AddContentController struct {
	rpcClient *rpc.Client
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
	if c.rpcClient == nil{
		c.rpcClient = rpcInit()
	}
	if islogin == nil {
		c.Data["isLogin"] = 0
		c.Redirect("/index", 302)
	} else {
		var content string
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")
		err := c.Ctx.Input.Bind(&content, "content")
		if err != nil {
			log.Fatalln("content nil: ", err)
		}
		hehe := addContent(content, c.rpcClient)
		c.Ctx.WriteString(hehe)
	}
}


