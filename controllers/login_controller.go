package controllers

import (
	"github.com/ulricqin/beego-blog/g"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	this.TplName = "login/login.html"
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

	this.Ctx.SetCookie("bb_name", name, 2592000, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "bb_password="+password+"; Max-Age=2592000; Path=/; httponly")

	this.Ctx.WriteString("")
}

func (this *LoginController) Logout() {
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "bb_name=0; Max-Age=0; Path=/;")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "bb_password=0; Max-Age=0; Path=/; httponly")
	this.Ctx.WriteString("logout")
}
