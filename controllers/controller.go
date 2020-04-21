package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	fmt.Print("RRRRRRRRRRRRRRRR")
	c.TplName = "index.html"
}
