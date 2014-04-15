package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/beego-blog/g"
	"github.com/ulricqin/goutils/paginator"
	"strconv"
)

type Checker interface {
	CheckLogin()
}

type BaseController struct {
	beego.Controller
	IsAdmin bool
}

func (this *BaseController) Prepare() {
	this.Data["BlogLogo"] = g.BlogLogo
	this.Data["BlogTitle"] = g.BlogTitle
	this.Data["BlogResume"] = g.BlogResume
	this.Data["RootName"] = g.RootName
	this.Data["RootEmail"] = g.RootEmail
	this.Data["RootPortrait"] = g.RootPortrait
	this.AssignIsAdmin()
	if app, ok := this.AppController.(Checker); ok {
		app.CheckLogin()
	}
}

func (this *BaseController) AssignIsAdmin() {
	bb_name := this.Ctx.GetCookie("bb_name")
	bb_password := this.Ctx.GetCookie("bb_password")
	if bb_name == "" || bb_password == "" {
		this.IsAdmin = false
		return
	}

	if bb_name != g.RootName || bb_password != g.RootPass {
		this.IsAdmin = false
	}

	this.IsAdmin = true
	this.Data["IsAdmin"] = this.IsAdmin
}

func (this *BaseController) SetPaginator(per int, nums int64) *paginator.Paginator {
	p := paginator.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *BaseController) GetIntWithDefault(paramKey string, defaultVal int) int {
	valStr := this.GetString(paramKey)
	var val int
	if valStr == "" {
		val = defaultVal
	} else {
		var err error
		val, err = strconv.Atoi(valStr)
		if err != nil {
			val = defaultVal
		}
	}
	return val
}

func (this *BaseController) JsStorage(action, key string, values ...string) {
	value := action + ":::" + key
	if len(values) > 0 {
		value += ":::" + values[0]
	}
	this.Ctx.SetCookie("JsStorage", value, 1<<31-1, "/")
}
