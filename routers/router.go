package routers

import (
	"beeTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/regist", &controllers.RegistController{})
}
