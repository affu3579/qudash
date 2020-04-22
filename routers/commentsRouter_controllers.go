package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:SessionController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:SessionController"],
        beego.ControllerComments{
            Method: "Home",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:SessionController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:SessionController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:SessionController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:SessionController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/zillani/qudash/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
