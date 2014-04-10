package controllers

type CatalogController struct {
	BaseController
}

func (this *CatalogController) Add() {
	this.Data["IsAdd"] = true
	this.Layout = "layout/admin.html"
	this.TplNames = "catalog/add.html"
}
