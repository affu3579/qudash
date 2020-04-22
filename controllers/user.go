package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/zillani/qudash/models"
	"log"
	"strconv"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Id        int    `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Mobile    int    `json:"Mobile"`
	UserName  string `json:"UserName"`
	Orders    int    `json:"Orders"`
	Password  string `json:"Password"`
	UserRole  uint8  `json:"UserRole"`
	IsActive  uint8  `json:"IsActive"`
}

type Users []User

var users []User

func init() {
	models.Connect()
}

// @Title Create User
// @Description create user
// @Param	body		body 	models.User	true		"The object content"
// @Success 200 {string} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Post() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	id, _ := models.CreateUser(user)
	this.Data["json"] = map[string]int64{"id": id}
	this.ServeJSON()
}

// @Title Get
// @Description find user by id
// @Param	id		path 	string	true		"the id you want to get"
// @Success 200 {user} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (this *UserController) Get() {
	id := this.Ctx.Input.Param(":id")
	uid, _ := strconv.Atoi(id)
	this.Ctx.Input.Bind(&id, "id")
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = models.GetUser(uid)
	this.ServeJSON()
}

// @Title GetAll
// @Description Get all users
// @Success 200 {object} models.User
// @Failure 500 server error
// @router / [get]
func (this *UserController) GetAll() {
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"], _ = models.GetUsers()
	this.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"The body"
// @Success 200 {user} models.User
// @Failure 403 :id is empty
// @router / [put]
func (this *UserController) Put() {
	userId:= this.Ctx.Input.Param(":id")
	log.Print("going to update id", userId)
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	uid ,_ := strconv.Atoi(userId)
	this.Ctx.Input.Bind(&uid, "id")
	log.Print("going to update id", userId)
	id, _ := models.UpdateUser(uid, user)
	this.Data["json"] = map[string]int{"id": id}
	this.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *UserController) Delete() {
	id := this.Ctx.Input.Param(":id")
	uid, _ := strconv.Atoi(id)
	this.Ctx.Input.Bind(&id, "id")
	uid , _ = models.DeleteUser(uid)
	this.Data["json"] = map[string]int{"id ": uid}
	this.ServeJSON()
}

func (this *UserController) Dashboard() {
	this.Data["users"], _ = models.GetUsers()
	this.TplName = "dashboard.html"
}
