package main

import (
	_ "./routers"

	"github.com/astaxie/beego"
	_ "github.com/garyburd/redigo/redis"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
