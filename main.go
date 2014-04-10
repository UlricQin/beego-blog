package main

import (
	"github.com/astaxie/beego"
	"github.com/ulricqin/beego-blog/g"
	_ "github.com/ulricqin/beego-blog/routers"
)

func main() {
	g.InitEnv()
	beego.Run()
}
