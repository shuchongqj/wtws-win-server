package wtws_mongodb

import (
	"context"
	"github.com/astaxie/beego/logs"
	"time"
	"wtws-server/models"
)

type UserLoginLog struct {
	UserId    int    `json:"userId"`
	LoginName string `json:"loginName"`
	WorkNo    string `json:"workNo"`
	LoginTime string `json:"loginTime"`
}

func (c *UserLoginLog) InitTable() string {
	return "user_login_log"
}

// InsertUserLoginLog 插入用户登录日志
func InsertUserLoginLog(userId int, loginName string, workNo string) error {
	loginLog := UserLoginLog{
		UserId:    userId,
		LoginName: loginName,
		WorkNo:    workNo,
		LoginTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	if _, err := models.MongodbClientDB.Collection(loginLog.InitTable()).InsertOne(context.TODO(), &loginLog); err != nil {
		logs.Error("[mongodb] 插入用户登录记录数据失败,失败信息:", err.Error())
		return err
	}
	return nil
}
