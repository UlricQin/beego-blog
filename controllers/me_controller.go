package controllers

type MeController struct {
	AdminController
}

func (this *MeController) Default() {
	this.TplNames = "me/default.html"
}
