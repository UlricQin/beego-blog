package controllers

import (
	"github.com/ulricqin/beego-blog/models"
	"github.com/ulricqin/beego-blog/models/blog"
	"github.com/ulricqin/beego-blog/models/catalog"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Read() {
	id, _ := this.GetInt(":id")
	this.Data["id"] = id
	// views + 1
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

	cp := catalog.OneById(int64(catalog_id))
	if cp == nil {
		this.Ctx.WriteString("catalog_id not exists")
		return
	}

	b := &models.Blog{Ident: ident, Title: title, Keywords: keywords, CatalogId: int64(catalog_id), Type: int8(aType), Status: int8(status)}
	_, err := blog.Save(b, content)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Redirect("/catalog/"+cp.Ident, 302)

}
