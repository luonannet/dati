package main

import (
	_ "dati/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.Listen.Graceful = true
	beego.BConfig.ServerName = "roland"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 36000
	beego.Run()
}
