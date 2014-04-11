package controllers

import (
	"fmt"
	"github.com/ulricqin/beego-blog/g"
	"github.com/ulricqin/beego-blog/models"
	"github.com/ulricqin/beego-blog/models/catalog"
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
	resume := this.GetString("resume")
	display_order := this.GetIntWithDefault("display_order", 0)

	if name == "" {
		this.Ctx.WriteString("name is blank")
		return
	}

	if ident == "" {
		this.Ctx.WriteString("ident is blank")
		return
	}

	_, header, err := this.GetFile("img")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	fmt.Println(header.Filename)

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
	o := &models.Catalog{Ident: ident, Name: name, Resume: resume, DisplayOrder: display_order, ImgUrl: "/" + imgPath}
	_, err = catalog.Save(o)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		filetool.Unlink(imgPath)
		return
	}

	this.Redirect("/", 302)
}
