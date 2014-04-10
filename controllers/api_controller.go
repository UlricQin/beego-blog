package controllers

import (
	"github.com/ulricqin/goutils/strtool"
	"github.com/ulricqin/beego-blog/models/catalog"
	"strings"
	"fmt"
)

type ApiController struct {
	BaseController
}

func (this *ApiController) Health() {
	fmt.Println(catalog.All()[0])
	this.Ctx.WriteString("ok")
}

func (this *ApiController) Md5() {
	p := this.GetString("p")
	this.Ctx.WriteString(strtool.Md5(strings.TrimSpace(p)))
}
