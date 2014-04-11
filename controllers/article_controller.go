package controllers

import (
	"github.com/ulricqin/beego-blog/models"
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
	title := this.GetString("title")
	ident := this.GetString("ident")
	keywords := this.GetString("keywords")
	catalog_id := this.GetIntWithDefault("catalog_id", -1)
	aType := this.GetIntWithDefault("type", -1)
	status := this.GetIntWithDefault("status", -1)
	content := this.GetString("content")

	if catalog_id == -1 || aType == -1 || status == -1 {
		this.Ctx.WriteString("catalog || type || status is illegal")
		return
	}

	// blog content save first
	// blog := &models.Blog{Ident:ident, Title:title, Keywords:keywords, CatalogId:catalog_id, }

	// add success
	// clear cache catalog/$id/article_ids
	// redirect to /catalog/xxxx
}
