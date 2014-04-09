package controllers

import (
	"github.com/ulricqin/goutils/strtool"
	"strings"
)

type ApiController struct {
	BaseController
}

func (this *ApiController) Health() {
	this.Ctx.WriteString("ok")
}

func (this *ApiController) Md5() {
	p := this.GetString("p")
	this.Ctx.WriteString(strtool.Md5(strings.TrimSpace(p)))
}
