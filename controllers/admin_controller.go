package controllers

import (
	"github.com/ulricqin/beego-blog/g"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) CheckLogin() {
	bb_name := this.Ctx.GetCookie("bb_name")
	bb_password := this.Ctx.GetCookie("bb_password")
	if bb_name == "" || bb_password == "" {
		this.Redirect("/login", 302)
		return
	}

	if bb_name != g.RootName || bb_password != g.RootPass {
		this.Redirect("/login", 302)
		return
	}

	this.IsAdmin = true
}
