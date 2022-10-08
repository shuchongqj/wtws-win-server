package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"wtws-server/common"
	"wtws-server/conf"
	"wtws-server/controllers"
)

func init() {

	version := conf.VERSION

	common.CorsDomain()
	//token认证，并解析userID
	//beego.InsertFilter(fmt.Sprintf("/%s/*", version), beego.BeforeRouter, common.Token)

	ns := beego.NewNamespace(fmt.Sprintf("/%s", version),
		beego.NSBefore(common.Auth),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),

		beego.NSNamespace("/station",
			beego.NSInclude(
				&controllers.StationController{},
			),
		),

		beego.NSNamespace("/enterprise",
			beego.NSInclude(
				&controllers.EnterpriseController{},
			),
		),

		beego.NSNamespace("/receive",
			beego.NSInclude(
				&controllers.ReceiveController{},
			),
		),

		beego.NSNamespace("/origin",
			beego.NSInclude(
				&controllers.OriginController{},
			),
		),

		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.GoodsController{},
			),
		),

		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),

		beego.NSNamespace("/driver",
			beego.NSInclude(
				&controllers.DriverController{},
			),
		),

		beego.NSNamespace("/cargoto",
			beego.NSInclude(
				&controllers.CargotoController{},
			),
		),

		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
		beego.NSNamespace("/truck-order",
			beego.NSInclude(
				&controllers.TruckOrderController{},
			),
		),
		beego.NSNamespace("/weigh-order",
			beego.NSInclude(
				&controllers.WeighOrderController{},
			),
		),
		beego.NSNamespace("/analysis",
			beego.NSInclude(
				&controllers.AnalysisController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
