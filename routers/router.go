package routers

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/beego-blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.AutoRouter(&controllers.ApiController{})

	beego.Router("/article/:id:int", &controllers.ArticleController{}, "get:Read")

	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/login", &controllers.LoginController{}, "post:DoLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/me", &controllers.MeController{}, "get:Default")
	beego.Router("/me/catalog/list", &controllers.CatalogController{}, "get:List")
}
