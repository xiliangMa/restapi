package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["restapi/controllers:ClusterController"] = append(beego.GlobalControllerRouter["restapi/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "ClusterList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:ClusterController"] = append(beego.GlobalControllerRouter["restapi/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "DeleteCluster",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:ClusterController"] = append(beego.GlobalControllerRouter["restapi/controllers:ClusterController"],
        beego.ControllerComments{
            Method: "AddCluster",
            Router: `/addCluster`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:HostController"] = append(beego.GlobalControllerRouter["restapi/controllers:HostController"],
        beego.ControllerComments{
            Method: "HostList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:HostController"] = append(beego.GlobalControllerRouter["restapi/controllers:HostController"],
        beego.ControllerComments{
            Method: "DeleteHost",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:HostController"] = append(beego.GlobalControllerRouter["restapi/controllers:HostController"],
        beego.ControllerComments{
            Method: "AddHost",
            Router: `/addhost`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:RancherServerController"] = append(beego.GlobalControllerRouter["restapi/controllers:RancherServerController"],
        beego.ControllerComments{
            Method: "RancherServerList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:RancherServerController"] = append(beego.GlobalControllerRouter["restapi/controllers:RancherServerController"],
        beego.ControllerComments{
            Method: "DeleteRancherServer",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:RancherServerController"] = append(beego.GlobalControllerRouter["restapi/controllers:RancherServerController"],
        beego.ControllerComments{
            Method: "AddRancherServer",
            Router: `/addRancherServer`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["restapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserList",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["restapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["restapi/controllers:UserController"] = append(beego.GlobalControllerRouter["restapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/addUser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
