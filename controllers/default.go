package controllers

import (
	_ "beeTest/models"
	"github.com/astaxie/beego"
	"github.com/go-ego/riot/types"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
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
		c.rpcInit()
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
		c.Data["content"] = search(content, c.rpcClient)
		c.TplName = "result.html"
	}
}

func (c *SearchController) rpcInit() {
	var err error
	c.rpcClient, err = jsonrpc.Dial("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
}

func search(content string, rpcClient *rpc.Client) types.SearchResp {
	req := SearchRequest{content}
	var res SearchResponse
	err := rpcClient.Call("RPCEngine.Search", req, &res)
	if err != nil {
		log.Fatalln("search error: ", err)
	}
	return res.Content
}

/*************  SEARCH  ******************/
type InterestingController struct {
	beego.Controller
}

func (c *InterestingController) Get() {
	c.TplName = "relax.html"
}
