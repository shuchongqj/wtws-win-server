package wtws_mongodb

import (
	"context"
	"github.com/astaxie/beego/logs"
	"time"
	"wtws-server/models"
)

type ApiResponseLoginLog struct {
	Path         string `json:"path"`
	ResponseCode string `json:"response_code"`
	ResponseBody string `json:"response_body"`
	ResPonseTime string `json:"response_time"`
}

func (c *ApiResponseLoginLog) InitTable() string {
	return "api_response_log_" + c.Path
}

// InsertApiLoginLog 插入接口访问日志
func InsertApiResponseLoginLog(path, userAgent, requestIp string) error {
	apiLog := ApiResponseLoginLog{
		Path:         path,
		ResPonseTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	if _, err := models.MongodbClientDB.Collection(apiLog.InitTable()).InsertOne(context.TODO(), &apiLog); err != nil {
		logs.Error("[mongodb] 插入接口访问日志失败,失败信息:", err.Error())
		return err
	}
	return nil
}
