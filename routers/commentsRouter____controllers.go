package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["wtws-server/controllers:AnalysisController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:AnalysisController"],
        beego.ControllerComments{
            Method: "GetAnalysisDetail",
            Router: "/detail",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:AnalysisController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:AnalysisController"],
        beego.ControllerComments{
            Method: "GetAnalysisOrderTypeDetail",
            Router: "/order-type-detail",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"],
        beego.ControllerComments{
            Method: "AddCargoto",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"],
        beego.ControllerComments{
            Method: "DeleteCargoto",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"],
        beego.ControllerComments{
            Method: "UpdateCargoto",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"],
        beego.ControllerComments{
            Method: "GetAllCargoto",
            Router: "/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:CargotoController"],
        beego.ControllerComments{
            Method: "GetCargotoList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:CategoryController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetAllCategory",
            Router: "/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:DriverController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:DriverController"],
        beego.ControllerComments{
            Method: "AddDriver",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:DriverController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:DriverController"],
        beego.ControllerComments{
            Method: "DeleteDriver",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:DriverController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:DriverController"],
        beego.ControllerComments{
            Method: "UpdateDriver",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:DriverController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:DriverController"],
        beego.ControllerComments{
            Method: "GetAllDriver",
            Router: "/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:DriverController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:DriverController"],
        beego.ControllerComments{
            Method: "GetDriverList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:DriverController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:DriverController"],
        beego.ControllerComments{
            Method: "TruckOrderList",
            Router: "/truck-order/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:EnterpriseController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:EnterpriseController"],
        beego.ControllerComments{
            Method: "GetUserEnterprises",
            Router: "/user/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "AddGoods",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "DeleteGoods",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "UpdateGoods",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetAllGoodsList",
            Router: "/all-list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:GoodsController"],
        beego.ControllerComments{
            Method: "GetGoodsList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "AddOrder",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "DeleteOrder",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "UpdateOrder",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "CheckOrder",
            Router: "/check",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetAllCheckedOrder",
            Router: "/checked",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "DownAllOrder",
            Router: "/down-all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetOrderInfo",
            Router: "/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "InvalidOrder",
            Router: "/invalid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OrderController"],
        beego.ControllerComments{
            Method: "GetOrderList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OriginController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OriginController"],
        beego.ControllerComments{
            Method: "AddOrigin",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OriginController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OriginController"],
        beego.ControllerComments{
            Method: "DeleteOrigin",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OriginController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OriginController"],
        beego.ControllerComments{
            Method: "UpdateOrigin",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OriginController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OriginController"],
        beego.ControllerComments{
            Method: "GetAllOriginList",
            Router: "/all-list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:OriginController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:OriginController"],
        beego.ControllerComments{
            Method: "GetOriginList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"],
        beego.ControllerComments{
            Method: "AddReceive",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"],
        beego.ControllerComments{
            Method: "DeleteReceive",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"],
        beego.ControllerComments{
            Method: "UpdateReceive",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"],
        beego.ControllerComments{
            Method: "GetAllReceiveList",
            Router: "/all-list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:ReceiveController"],
        beego.ControllerComments{
            Method: "GetReceiveList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "DeleteRole",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "AddRole",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetAllRoleList",
            Router: "/all-list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetRoleFunctions",
            Router: "/functions",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:RoleController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:RoleController"],
        beego.ControllerComments{
            Method: "AddRoleFunctions",
            Router: "/functions",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:StationController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:StationController"],
        beego.ControllerComments{
            Method: "UpdateStation",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:StationController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:StationController"],
        beego.ControllerComments{
            Method: "AddStation",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:StationController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:StationController"],
        beego.ControllerComments{
            Method: "DeleteStation",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:StationController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:StationController"],
        beego.ControllerComments{
            Method: "GetStationList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:StationController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:StationController"],
        beego.ControllerComments{
            Method: "GetUserStation",
            Router: "/user/list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "AddTruckOrder",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "DeleteTruckOrder",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "UpdateTruckOrder",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "CheckTruckOrder",
            Router: "/check",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "DownAllTruckOrder",
            Router: "/down-all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "InvalidTruckOrder",
            Router: "/invalid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "GetTruckOrderList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "AddSentDirectTruckOrder",
            Router: "/sent-direct",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:TruckOrderController"],
        beego.ControllerComments{
            Method: "GetTruckOrderByVehicle",
            Router: "/vehicle",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUserInfo",
            Router: "/add",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserLIst",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUserLoginName",
            Router: "/login-name",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "LogOut",
            Router: "/logout",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "RestUserPwd",
            Router: "/reset-pwd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUserRole",
            Router: "/role",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdatePwd",
            Router: "/update-pwd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:UserController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUserInfo",
            Router: "/user-info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "AddWeighOrder",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "DeleteWeighOrder",
            Router: "/",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "UpdateWeighOrder",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "DownAllWeighOrder",
            Router: "/down-all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "FinishWeighOrder",
            Router: "/finish",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "InvalidWeighOrder",
            Router: "/invalid",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "GetWeighOrderList",
            Router: "/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "ScanVehicle",
            Router: "/scan-vehicle",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "TareWight",
            Router: "/tare-wight",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "WaitFinishOrder",
            Router: "/wait-finish",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "CheckWareHouseGoods",
            Router: "/ware-house-check-goods",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"] = append(beego.GlobalControllerRouter["wtws-server/controllers:WeighOrderController"],
        beego.ControllerComments{
            Method: "WarehouseCheckWeighOrder",
            Router: "/warehouse-check",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
