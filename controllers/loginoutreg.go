package controllers

import (
	"beeTest/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*************  LOGIN  ******************/
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	o := orm.NewOrm()

	user := models.User{}
	user.Username = username

	err := o.Read(&user, "Username")
	if err != nil {
		c.Ctx.WriteString("用户不存在")
	} else if user.Password != password {
		c.Ctx.WriteString("密码不对")
	} else {
		c.Ctx.SetCookie("islogin", "1", 3600)
		c.SetSession("username", username)
		c.Redirect("/index", 302)
	}

}

/*************  LOGOUT  ******************/
type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	c.Ctx.SetCookie("islogin", "0", 3600)
	if c.GetSession("username") != nil {
		c.DelSession("username")
		c.DestroySession()
	}
	c.Redirect("/index", 302)
}

/*************  REGIST  ******************/
type RegistController struct {
	beego.Controller
}

func (c *RegistController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	o := orm.NewOrm()

	user := models.User{}

	user.Username = username
	user.Password = password

	_, err := o.Insert(&user)
	if err != nil {
		c.Redirect("/index", 302)
		return
	} else {
		c.Redirect("/index", 302)
	}
	//c.TplName = "search.html"
}
