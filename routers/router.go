// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/zillani/qudash/controllers"
	"github.com/zillani/qudash/filters"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/chat", &controllers.ChatController{})
	beego.Router("/join", &controllers.ChatController{}, "post:Join")
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.SessionController{}, "get:Home")
	beego.Router("/login", &controllers.SessionController{}, "get:Login")
	beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
	beego.Router("/dashboard", &controllers.UserController{}, "get:Dashboard")
	beego.Router("/getFromCache", &controllers.CacheController{}, "get:GetFromCache")
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LogManager)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
