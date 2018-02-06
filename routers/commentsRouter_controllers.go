package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/answer/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetAllAnswer",
			Router: `/answer/getAll`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "CreateAdmin",
			Router: `/createAdmin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "CreateBangdan",
			Router: `/createBangdan`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "GetBangdan",
			Router: `/getBangdan`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/inputAnswer`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/question/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:AdminController"] = append(beego.GlobalControllerRouter["dati/controllers:AdminController"],
		beego.ControllerComments{
			Method: "SetQuestionID",
			Router: `/setQuestionID`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:UserController"] = append(beego.GlobalControllerRouter["dati/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/answer`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:UserController"] = append(beego.GlobalControllerRouter["dati/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/answer/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:UserController"] = append(beego.GlobalControllerRouter["dati/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getMyAllAnswer`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:UserController"] = append(beego.GlobalControllerRouter["dati/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["dati/controllers:UserController"] = append(beego.GlobalControllerRouter["dati/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
