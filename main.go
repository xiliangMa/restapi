package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xiliangMa/restapi/routers"
	_ "github.com/xiliangMa/restapi/sysinit"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
