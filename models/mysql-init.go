package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"wtws-server/conf"
)

func InitORM() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	userName := conf.DATABASE_USERNAME
	password := conf.DATABASE_PASSWORD
	dbBaseUrl := conf.DATABASE_URL
	dbPort := conf.DATABASE_PORT
	dbName := conf.DATABASE_NAME
	logs.Info("[mysql]\t", userName+":"+password+"@tcp("+dbBaseUrl+":"+dbPort+")/"+dbName)
	if registDataBaseErr := orm.RegisterDataBase("default", "mysql",
		userName+":"+password+"@tcp("+dbBaseUrl+":"+dbPort+")/"+dbName+"?charset=utf8mb4"); registDataBaseErr != nil {
		logs.Error("[mysql]  注册连接数据库失败，失败信息：", registDataBaseErr.Error())
		panic(registDataBaseErr)
	} else {
		logs.Info("[mysql]  连接成功")
	}
}
