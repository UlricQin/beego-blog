package controllers

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Read() {
	id, _ := this.GetInt(":id")
	this.Data["id"] = id
	this.TplNames = "blog/read.html"
}



