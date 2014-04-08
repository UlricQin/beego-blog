package routers

import (
	"github.com/ulricqin/beego-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.ApiController{})

	beego.Router("/article/:id:int", &controllers.ArticleController{}, "get:Read")

	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/login", &controllers.LoginController{}, "post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

}
