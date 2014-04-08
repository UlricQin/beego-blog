package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/beego-blog/g"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	this.TplNames = "login/login.html"
}

func (this *LoginController) DoLogin() {
	name := this.GetString("name")
	if name == "" {
		this.Ctx.WriteString("name is blank")
		return
	}
	password := this.GetString("password")
	if password == "" {
		this.Ctx.WriteString("password is blank")
		return
	}

	if g.RootName != name {
		this.Ctx.WriteString("name is incorrect")
		return
	}

	if g.RootPass != password {
		this.Ctx.WriteString("password is incorrect")
		return
	}

	//TODO: write cookie

	this.Ctx.WriteString("")
}

func (this *LoginController) Logout() {
	this.Ctx.WriteString("logout")
}




