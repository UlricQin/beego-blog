package main

import (
	_ "github.com/ulricqin/beego-blog/routers"
	"github.com/ulricqin/beego-blog/g"
	"github.com/astaxie/beego"
)

func main() {
	g.InitEnv()
	beego.Run()
}

