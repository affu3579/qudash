package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/beego/i18n"
	"github.com/zillani/qudash/controllers"
	_ "github.com/zillani/qudash/routers"
)

func main() {
	fmt.Print("INSIDE MAIN")
	beego.ErrorController(&controllers.ErrorController{})
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
