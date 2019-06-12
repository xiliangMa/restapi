package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Authorize",
            Router: `/authorize`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Signin",
            Router: `/signin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "ClusterList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "DeleteCluster",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:ClusterController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "AddCluster",
            Router: `/addCluster`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:HostController"],
        beego.ControllerComments{
            Method: "HostList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:HostController"],
        beego.ControllerComments{
            Method: "DeleteHost",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:HostController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:HostController"],
        beego.ControllerComments{
            Method: "AddHost",
            Router: `/addhost`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderDetailController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderDetailController"],
        beego.ControllerComments{
            Method: "OrderDetailList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderDetailController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderDetailController"],
        beego.ControllerComments{
            Method: "DeleteOrderDetail",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderDetailController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderDetailController"],
        beego.ControllerComments{
            Method: "AddOrderDetail",
            Router: `/addOrderDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderMasterController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderMasterController"],
        beego.ControllerComments{
            Method: "OrderMasterList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderMasterController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderMasterController"],
        beego.ControllerComments{
            Method: "DeleteOrderMaster",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderMasterController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderMasterController"],
        beego.ControllerComments{
            Method: "AddOrderMaster",
            Router: `/addOrderMaster`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderRenewalController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderRenewalController"],
        beego.ControllerComments{
            Method: "OrderRenewalList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderRenewalController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderRenewalController"],
        beego.ControllerComments{
            Method: "DeleteOrderRenewal",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderRenewalController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:OrderRenewalController"],
        beego.ControllerComments{
            Method: "AddOrderRenewal",
            Router: `/addOrderRenewal`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:PromotionController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "PromotionList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:PromotionController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "DeletePromotion",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:PromotionController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "AddPromotion",
            Router: `/addPromotion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:RancherServerController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:RancherServerController"],
        beego.ControllerComments{
            Method: "RancherServerList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:RancherServerController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:RancherServerController"],
        beego.ControllerComments{
            Method: "DeleteRancherServer",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:RancherServerController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:RancherServerController"],
        beego.ControllerComments{
            Method: "AddRancherServer",
            Router: `/addRancherServer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/addUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/xiliangMa/restapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserInfo",
            Router: `/info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
