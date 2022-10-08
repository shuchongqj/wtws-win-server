package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
	"wtws-server/conf"
	"wtws-server/models"
	_ "wtws-server/routers"
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("[server]  service start failed!!!! connect database failed!!!!!!")
		}
	}()
	conf.InitService()
	models.InitORM()
	models.InitRedisClient()
	models.InitMongoDBClient()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("[server]  service start failed!!!!")
		}
	}()
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	orm.Debug = true
	beego.Run()
}
