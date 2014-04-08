package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Read() {
	id, _ := this.GetInt(":id")
	this.Data["id"] = id
	this.TplNames = "blog/read.html"
}



