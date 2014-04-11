package controllers

import (
	"github.com/ulricqin/beego-blog/models/catalog"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Read() {
	id, _ := this.GetInt(":id")
	this.Data["id"] = id
	this.TplNames = "article/read.html"
}

func (this *ArticleController) Add() {
	this.Data["Catalogs"] = catalog.All()
	this.Data["IsPost"] = true
	this.Layout = "layout/admin.html"
	this.TplNames = "article/add.html"
}

func (this *ArticleController) DoAdd() {
}
