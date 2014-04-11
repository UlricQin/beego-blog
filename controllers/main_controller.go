package controllers

import (
	"github.com/ulricqin/beego-blog/models/catalog"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Catalogs"] = catalog.All()
	this.Data["PageTitle"] = "首页"
	this.Layout = "layout/default.html"
	this.TplNames = "index.html"
}
