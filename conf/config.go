package conf

import "github.com/astaxie/beego"

const (
	VERSION           = "v1"
	SECREKEY          = "sunrise"
	DATABASE_PORT     = "33060"
	DATABASE_NAME     = "wtws-win"
	DATABASE_USERNAME = "root"
	DATABASE_PASSWORD = "root"
	MONGODB_SOURCE    = "wtws"
	REDIS_PORT        = "6379"
	ADMIN_ROLE_ID     = "1"
	DATABASE_URL      = "192.168.31.179"
	REDIS_IP          = "192.168.31.179"
	MONGODB_URI       = "mongodb://192.168.31.179:27017"
)

const (
	appname             = "wtws-server"
	httpPort            = 19009
	runmode             = "dev"
	copyrequestbody     = true
	enableDocs          = true
	routerCaseSensitive = false
	recoverPanic        = true
	enableAdmin         = true
	adminAddr           = "localhost"
	adminPort           = 12100
)

func InitService() {

	beego.BConfig.CopyRequestBody = copyrequestbody
	beego.BConfig.AppName = appname
	beego.BConfig.RunMode = runmode
	beego.BConfig.RecoverPanic = recoverPanic
	beego.BConfig.RouterCaseSensitive = routerCaseSensitive
	beego.BConfig.WebConfig.EnableDocs = enableDocs

	beego.BConfig.Listen.HTTPPort = httpPort
	beego.BConfig.Listen.EnableAdmin = enableAdmin
	beego.BConfig.Listen.AdminAddr = adminAddr
	beego.BConfig.Listen.AdminPort = adminPort
}
