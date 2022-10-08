package wtws_mongodb

import (
	"context"
	"github.com/astaxie/beego/logs"
	"strings"
	"time"
	"wtws-server/models"
)

type ApiRequestLoginLog struct {
	Path          string `json:"path"`
	UserID        int    `json:"user_id"`
	UserName      string `json:"user_name"`
	RequestMethod string `json:"request_method"`
	RequestBody   string `json:"request_body"`
	UserAgent     string `json:"user_agent"`
	RequestTime   string `json:"request_time"`
	RequestIP     string `json:"request_ip"`
}

func (c *ApiRequestLoginLog) InitTable() string {
	pathStr := c.Path
	pathStr = strings.Replace(pathStr, "/", "_", -1)
	pathStr = strings.Replace(pathStr, "-", "_", -1)
	return "api_request_log" + pathStr
}

// InsertApiLoginLog 插入接口访问日志
func InsertApiRequestLoginLog(path, requestMethod, requestBody, userAgent, requestIp, userName string, userID int) error {
	apiLog := ApiRequestLoginLog{
		Path:          path,
		UserID:        userID,
		UserName:      userName,
		RequestMethod: requestMethod,
		RequestBody:   requestBody,
		UserAgent:     userAgent,
		RequestIP:     requestIp,
		RequestTime:   time.Now().Format("2006-01-02 15:04:05"),
	}
	if _, err := models.MongodbClientDB.Collection(apiLog.InitTable()).InsertOne(context.TODO(), &apiLog); err != nil {
		logs.Error("[mongodb] 插入接口访问日志失败,失败信息:", err.Error())
		return err
	}
	return nil
}
