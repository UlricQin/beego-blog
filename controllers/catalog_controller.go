package controllers

import (
	"fmt"
	"github.com/ulricqin/beego-blog/g"
	"github.com/ulricqin/goutils/filetool"
	"time"
)

type CatalogController struct {
	AdminController
}

func (this *CatalogController) Add() {
	this.Data["IsAdd"] = true
	this.Layout = "layout/admin.html"
	this.TplNames = "catalog/add.html"
}

func (this *CatalogController) DoAdd() {
	name := this.GetString("name")
	ident := this.GetString("ident")
	// resume := this.GetString("resume")
	// display_order := this.GetIntWithDefault("display_order", 0)

	if name == "" {
		this.Ctx.WriteString("name is blank")
		return
	}

	if ident == "" {
		this.Ctx.WriteString("ident is blank")
		return
	}

	// 检查ident是否已经存在

	_, header, err := this.GetFile("img")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	ext := filetool.Ext(header.Filename)

	catalogDir := "static/uploads/catalogs"
	filetool.InsureDir(catalogDir)
	imgPath := fmt.Sprintf("%s/%s_%d%s", catalogDir, ident, time.Now().Unix(), ext)

	if err = this.SaveToFile("img", imgPath); err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	if g.UseQiniu {
		// 上传到七牛，并且返回一个url
	}

	// 保存分类信息到DB

	this.Redirect("/", 302)

}
