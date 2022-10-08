package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	jwtInfo "github.com/wangcong0918/sunrise/utils/jwt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"wtws-server/conf"
	"wtws-server/models"
	wtws_mongodb "wtws-server/models/wtws-mongodb"
	wtws_mysql "wtws-server/models/wtws-mysql"
)

func IsPublicFunctions(path string) bool {

	isPublicFunc := false

	dataArr := wtws_mysql.GetAllSPublicFunction()
	for _, data := range dataArr {
		if strings.Contains(path, data.Path) {
			isPublicFunc = true
		}
	}

	return isPublicFunc
}

//备用  http web鉴权
//func TokenBK(ctx *context.Context) {
//	authUrl, _ := env.MustGet("PUBLICRERUESTURL")
//	authPort, _ := env.MustGet("PUBLICRERUESTPORT")
//	authPath, _ := env.MustGet("PUBLICRERUESTPATH")
//
//	logs.Debug("\t\t[request][url]==>", ctx.Request.RequestURI, "\t[method]=>", ctx.Request.Method, "\t[body]=>", ctx.Request.Body, "\t[time]==>", time.Now())
//	//请求认真服务，进行token认证，解析user信息
//	var authUrlPath = "http://" + authUrl + ":" + authPort + authPath
//	type emptyStruct struct {
//	}
//	payloadJson, _ := json.Marshal(emptyStruct{})
//	payload := strings.NewReader(string(payloadJson))
//	request, _ := http.NewRequest("POST", authUrlPath, payload)
//
//	authorization := ctx.Request.Header.Get("Authorization")
//	request.Header.Add("Authorization", authorization)
//	request.Header.Add("url", ctx.Request.RequestURI)
//
//	client := &http.Client{}
//	response, _ := client.Do(request)
//	defer response.Body.Close()
//	bodyBytes, _ := ioutil.ReadAll(response.Body)
//	var authResult third_party.AuthResponse
//	parseErr := json.Unmarshal(bodyBytes, &authResult)
//	if parseErr != nil || authResult.Code == -1 {
//		ctx.Output.JSON(ResponseStatus(-1, "", nil), false, false)
//	} else if authResult.Code == -3 {
//		ctx.Output.JSON(ResponseStatus(-3, "", nil), false, false)
//	} else {
//		ctx.Input.SetData("userId", authResult.Result.UserID)
//	}
//}

func Auth(ctx *context.Context) {
	requestUrl := strings.ToLower(ctx.Request.RequestURI)
	regex1 := regexp.MustCompile(`^(.*)\?.*`)
	requestUrlArr := regex1.FindStringSubmatch(requestUrl)
	if len(requestUrlArr) >= 1 {
		requestUrl = requestUrlArr[1]
	}
	requestMethod := strings.ToLower(ctx.Request.Method)

	reqBodyStr := ""

	if requestMethod != "get" {
		defer ctx.Request.Body.Close()
		reqBodyBytes, _ := io.ReadAll(ctx.Request.Body)
		reqBodyStr = string(reqBodyBytes)
		logs.Info("[service] request body 信息:", string(reqBodyBytes))

	}

	//判断是否是公共接口，如果是的话跳出鉴权
	isPublic := IsPublicFunctions(requestUrl)
	if isPublic {

		return
	}

	var user *wtws_mysql.SUser

	secretKey := conf.SECREKEY
	if token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		}); err == nil && token.Valid {
		userInfo, _ := jwtInfo.JwtParseUser(token.Raw)

		var userId int //mysql user info

		var getUserRedisKeyErr, getUserErr error

		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			var redisUserIDStr string
			if redisUserIDStr, getUserRedisKeyErr = models.RedisGet(userInfo.UserID); getUserRedisKeyErr != nil || len(redisUserIDStr) == 0 {
				logs.Error("[server]  鉴权失败，获取KeyUserID失败，失败信息:", getUserRedisKeyErr.Error())
			}
			wg.Done()
		}()

		go func() {
			if userId, getUserErr = strconv.Atoi(userInfo.UserID); getUserErr != nil {
				logs.Error("[server]  鉴权失败，转换UserID类型失败，失败信息:", getUserErr.Error())
			} else if user, getUserErr = wtws_mysql.GetSUserById(userId); getUserErr != nil || user == nil {
				logs.Error("[server]  鉴权失败，查询用户信息失败，失败信息:", getUserErr.Error())
			}
			wg.Done()
		}()

		wg.Wait()

		if getUserRedisKeyErr != nil || getUserErr != nil || user.Id == 0 {
			logs.Error("[service]  解析token信息失败\t请求的url:", requestUrl, "\t请求方式:", requestMethod)
			go func() {
				wtws_mongodb.InsertApiRequestLoginLog(requestUrl, requestMethod, reqBodyStr, ctx.Input.Header("User-Agent"), ctx.Input.Header("Origin"), "", 0)
			}()
			ctx.Output.JSON(ResponseStatus(-3, "", nil), false, false)
			return
		}

	} else {
		logs.Error("[service]  解析token信息失败\t请求的url:", requestUrl, "\t请求方式:", requestMethod)
		go func() {
			wtws_mongodb.InsertApiRequestLoginLog(requestUrl, requestMethod, reqBodyStr, ctx.Input.Header("User-Agent"), ctx.Input.Header("Origin"), "", 0)
		}()
		ctx.Output.JSON(ResponseStatus(-3, "", nil), false, false)
		return
	}

	if functions := wtws_mysql.GetFunctionByPath(requestUrl); len(functions) > 0 {
		if isAllow := wtws_mysql.CheckFunctionByUserID(user.Id, requestUrl, requestMethod); !isAllow {
			logs.Error("[service]  检查authorations中的userId是否存在失败。userID:", user.Id, "\t请求的url:", requestUrl, "\t请求方式:", requestMethod)
			go func() {
				wtws_mongodb.InsertApiRequestLoginLog(requestUrl, requestMethod, reqBodyStr, ctx.Input.Header("User-Agent"), ctx.Input.Header("Origin"), "", 0)
			}()
			ctx.Output.JSON(ResponseStatus(-3, "", nil), false, false)
			return
		}
	}

	go func() {
		wtws_mongodb.InsertApiRequestLoginLog(requestUrl, requestMethod, reqBodyStr, ctx.Input.Header("User-Agent"), ctx.Input.Header("Origin"), user.DisplayName, user.Id)
	}()
	ctx.Input.SetData(conf.CTX_CONTEXT_USER, user)
}

func CorsDomain() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

}
