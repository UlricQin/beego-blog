package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/beego-blog/g"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["RootName"] = g.RootName
	this.Data["RootEmail"] = g.RootEmail
	this.Data["RootPortrait"] = g.RootPortrait
	this.TplNames = "index.tpl"
}
