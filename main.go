package main

import (
	_ "./routers"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	// // set default database
	// orm.RegisterDataBase("default", "mysql", "root:cyz8023721@tcp(127.0.0.1:3307)/test?charset=utf8", 30)

	// // register model
	// orm.RegisterModel(new(User))

	// // create table
	// orm.RunSyncdb("default", false, true)
}
