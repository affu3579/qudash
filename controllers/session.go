package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

// @router / [get]
func (this *SessionController) Home() {
	isAuthenticated := this.GetSession("authenticated")
	fmt.Println(isAuthenticated)
	if isAuthenticated == nil || isAuthenticated == false {
		this.Ctx.WriteString("You are unauthorized to view the page.")
		return
	}
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.WriteString("Home Page")
}

// @router / [get]
func (this *SessionController) Login() {
	this.SetSession("authenticated", true)
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.WriteString("You have successfully logged in.")
}

// @router / [get]
func (this *SessionController) Logout() {
	this.SetSession("authenticated", false)
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Ctx.WriteString("You have successfully logged out.")
}
