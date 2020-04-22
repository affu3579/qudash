package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "Page Not Found"
	c.TplName = "404.html"
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "Internal Server Error"
	c.TplName = "500.html"
}

func (c *ErrorController) ErrorGeneric() {
	c.Data["content"] = "Some Error Occurred"
	c.TplName = "genericerror.html"
}
