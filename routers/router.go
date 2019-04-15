// @APIVersion 1.0.0
// @Title Backend Service API
// @Description Backend Service API
// @TermsOfServiceUrl http://beego.me/
// @Contact astaxie@gmail.com
// @License MIT License
// @LicenseUrl https://opensource.org/licenses/MIT
package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/xiliangMa/restapi/controllers"
	"github.com/xiliangMa/restapi/utils"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/hosts",
			beego.NSInclude(
				&controllers.HostController{},
			),
		),
		beego.NSNamespace("/clusters",
			beego.NSInclude(
				&controllers.ClusterController{},
			),
		),
		beego.NSNamespace("/rancherservers",
			beego.NSInclude(
				&controllers.RancherServerController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/promotion",
			beego.NSInclude(
				&controllers.PromotionController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
	)

	var isLogin = func(ctx *context.Context) {
		_, code := utils.CheckToken(ctx.Input.Header("token"))
		if code != 200 {
			ctx.Redirect(401, "/swagger")
		}
	}
	beego.InsertFilter("/v1/hosts/", beego.BeforeRouter, isLogin)



	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.AddNamespace(ns)
}
