package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"os"
	"time"
	"wtws-server/conf"
)

var RedisClient *redis.Client

func InitRedisClient() {
	redisAddr := fmt.Sprintf("%s:%s", conf.REDIS_IP, conf.REDIS_PORT)
	logs.Info("[redis]\t", redisAddr)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PWD"), // no password set
		DB:       0,                      // use default DB
	})
	if _, err := client.Ping().Result(); err != nil {
		logs.Error("[redis]  redis数据库连接失败. 失败信息:", err)
		panic(err)
	} else {
		logs.Info("[redis]  连接成功")
	}

	RedisClient = client
}

func RedisSet(key, value string) (err error) {
	if err = RedisClient.Set(key, value, 720*time.Hour).Err(); err != nil {
		logs.Error("[redis]  RedisSet error. error message:", err.Error())
	}
	return err
}
func RedisGet(key string) (result string, err error) {
	if result, err = RedisClient.Get(key).Result(); err != nil {
		logs.Error("[redis]  RedisGet error. error message:", err.Error())
		result = ""
	}
	return result, err
}

func RedisDel(key string) (err error) {
	if err = RedisClient.Del(key).Err(); err != nil {
		logs.Error("[redis]  RedisDel error. error message:", err.Error())
	} else {
		logs.Info("[redis]  RedisDel success. ")
	}
	return err
}

func RedisDelAndSet(key, value string) (err error) {
	if err = RedisClient.Del(key).Err(); err != nil {
		logs.Error("[redis] RedisDelSet  error.  error message:", err.Error())
	} else if err = RedisClient.Set(key, value, 720*time.Hour).Err(); err != nil {
		logs.Error("[redis] RedisDelSet  error.  error message:", err.Error())
	} else {
		logs.Info("[redis]  RedisDelSet  success")
	}
	return err
}
