package controllers

import (
	"fmt"
	"github.com/ulricqin/beego-blog/g"
	"github.com/ulricqin/beego-blog/models/catalog"
	"github.com/ulricqin/goutils/filetool"
	"github.com/ulricqin/goutils/strtool"
	"strings"
	"time"
)

const (
	EDITOR_IMG_DIR = "static/uploads/editor"
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

func (this *ApiController) Upload() {
	result := map[string]interface{}{
		"success": false,
	}

	defer func() {
		this.Data["json"] = &result
		this.ServeJson()
	}()

	_, header, err := this.GetFile("image")
	if err != nil {
		return
	}

	ext := filetool.Ext(header.Filename)
	imgPath := fmt.Sprintf("%s/%d%s", EDITOR_IMG_DIR, time.Now().Unix(), ext)

	filetool.InsureDir(EDITOR_IMG_DIR)
	err = this.SaveToFile("image", imgPath)
	if err != nil {
		return
	}

	imgUrl := "/" + imgPath
	if g.UseQiniu {
		if addr, er := g.UploadFile(imgPath, imgPath); er != nil {
			return
		} else {
			imgUrl = addr
			filetool.Unlink(imgPath)
		}
	}

	result["link"] = imgUrl
	result["success"] = true

}

func (this *ApiController) Markdown() {
	if this.IsAjax() {
		result := map[string]interface{}{
			"success": false,
		}
		action := this.GetString("action")
		switch action {
		case "preview":
			content := this.GetString("content")
			result["preview"] = g.RenderMarkdown(content)
			result["success"] = true
		}
		this.Data["json"] = result
		this.ServeJson()
	}
}
