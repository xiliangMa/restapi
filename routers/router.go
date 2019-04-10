// @APIVersion 1.0.0
// @Title Backend Service API
// @Description Backend Service API
// @TermsOfServiceUrl http://beego.me/
// @Contact astaxie@gmail.com
// @License MIT License
// @LicenseUrl https://opensource.org/licenses/MIT
package routers

import (
	"github.com/xiliangMa/restapi/controllers"

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
	)
	beego.AddNamespace(ns)
}
