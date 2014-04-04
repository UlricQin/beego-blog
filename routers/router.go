package routers

import (
	"github.com/ulricqin/beego-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
