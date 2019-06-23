package controllers

import (
	_ "beeTest/models"
	"beeTest/tools"
	_ "beeTest/tools"
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	"log"
	"strconv"
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
				c.Data["message"] = err
				c.TplName = "message.html"
				log.Fatalln("client err: ", err)
				return
			}
		}
		var content string
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")
		err := c.Ctx.Input.Bind(&content, "content")
		if err != nil {
			c.Data["message"] = err
			c.TplName = "message.html"
			log.Fatalln("client err: ", err)
			return
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
		if c.client == nil {
			err := c.clientInit()
			if err != nil {
				c.Data["message"] = err
				c.TplName = "message.html"
				log.Println("client err: ", err)
				return
			}
		}
		indicies := getIndexName(c.catIndices())
		c.Data["indicies"] = indicies
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
		if c.client == nil {
			err := c.clientInit()
			if err != nil {
				c.Data["message"] = err
				c.TplName = "message.html"
				log.Println("client err: ", err)
				return
			}
		}
		c.Data["isLogin"] = 1
		c.Data["username"] = c.GetSession("username")

		ins := map[string]interface{}{}
		fieldNum, err := strconv.Atoi(c.GetString("fieldNum"))
		if err != nil{
			c.Data["message"] = err
			c.TplName = "message.html"
			log.Println("client err: ", err)
			return
		}
		index := c.GetString("index")
		for i:=0; i<fieldNum; i++{
			suffix := strconv.Itoa(i)
			ins[c.GetString("field"+suffix)] = c.GetString("content" + suffix)
		}
		id, err := c.addIndex(ins, index, "")
		if err!=nil{
			c.Data["message"] = err
			c.TplName = "message.html"
			log.Println("client err: ", err)
			return
		}
		c.Data["id"] = id
		c.Data["message"] = "success"
		c.TplName = "message.html"
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

/*************  XIALA  ******************/
type XIALAController struct {
	beego.Controller
}

func (c *XIALAController) Get() {
	hehe := []string{"a", "b", "c"}
	c.Data["hehe"] = hehe
	c.TplName = "xiala.html"
}

/*************  fileTraverse  ******************/
type TraverseController struct {
	beego.Controller
}

func (c *TraverseController) Get() {
	err := tools.Traverse()
	if err == nil{
		c.Data["message"] = "success!!!"
	}else {
		c.Data["message"] = err
	}
	c.TplName = "message.html"
}
