package controllers

import (
	"beeTest/models"
	_ "beeTest/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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


/*************  SEARCH  ******************/
type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	//v := c.GetSession("loginuser")
	//jsoninfo := ""
	c.Data["isLogin"] = 0
	c.Data["username"] = "hehe"
	c.TplName = "search.html"
}

func (c *SearchController) Post() {
	v := c.GetSession("loginuser")
	islogin:=true
	if v == nil {
		islogin=false
	}
	if islogin{
		jsoninfo := c.GetString("content")
		c.Data["text"] = jsoninfo
		c.TplName = "search.html"
	}else{
		c.TplName = "alert.html"

	}
}


/*************  LOGIN  ******************/
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	username:=c.GetString("username")
	password:=c.GetString("password")

	o := orm.NewOrm()

	user := models.User{}
	user.Username = username

	err := o.Read(&user, "Username")
	if err != nil{
		c.Ctx.WriteString("用户不存在")
	}else if user.Password != password{
		c.Ctx.WriteString("密码不对")
	}else{
		c.Ctx.WriteString("yes")
	}

}


/*************  LOGOUT  ******************/
type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Post() {
	v := c.GetSession("loginuser")
	islogin:=false
	if v != nil {
		//删除指定的session
		c.DelSession("loginuser")
		//销毁全部的session
		c.DestroySession()
		islogin=true

	}
	c.Data["json"]=map[string]interface{}{"islogin":islogin};
	c.ServeJSON();
	c.TplName = "search.html"
}


/*************  REGIST  ******************/
type RegistController struct {
	beego.Controller
}

func (c *RegistController) Post() {
	username:=c.GetString("username")
	password:=c.GetString("password")

	o := orm.NewOrm()

	user := models.User{}
	user.Username = username
	user.Password = password

	_, err := o.Insert(&user)
	if err != nil{
		c.Ctx.WriteString("no")
		//c.Redirect("/index", 302)
		return
	}else{
		c.Ctx.WriteString("yes")
	}
	//c.TplName = "search.html"
}


