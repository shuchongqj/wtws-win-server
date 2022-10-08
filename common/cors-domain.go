package common

func CorsDomain2() {

	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	//AllowAllOrigins: true,
	//	AllowOrigins:  []string{"http://47.114.163.213:20000"},
	//	AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:  []string{"authorization", "access-token", "a-auth-token", "x-auth-token", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	ExposeHeaders: []string{"authorization", "access-token", "Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	//AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	//AllowMethods: []string{"*"},
	//	//AllowHeaders: []string{"*"},
	//	//AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	//ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	//ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))
	//beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
	//	var success = []byte("SUPPORT OPTIONS")
	//	origin := ctx.Input.Header("Origin")
	//	ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
	//	ctx.Output.Header("Access-Control-Max-Age", "3600")
	//	ctx.Output.Header("Access-Control-Allow-Headers", "X-Custom-Header,accept,Content-Type,Access-Token")
	//	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	//	ctx.Output.Header("Access-Control-Allow-Origin", origin)
	//	if ctx.Input.Method() == http.MethodOptions {
	//		// options请求，返回200
	//		ctx.Output.SetStatus(http.StatusOK)
	//		_ = ctx.Output.Body(success)
	//	}
	//})
}
