package controllers

type CatalogController struct {
	BaseController
}

func (this *CatalogController) List() {
	this.Layout = "layout/admin.html"
	this.TplNames = "catalog/list.html"
}
