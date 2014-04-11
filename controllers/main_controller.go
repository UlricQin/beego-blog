package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["PageTitle"] = "首页"
	this.Layout = "layout/default.html"
	this.TplNames = "index.html"
}
