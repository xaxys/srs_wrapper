package middleware

import (
	"fmt"
	"srs_wrapper/config"
	"srs_wrapper/model"

	"net/url"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

var (
	TcUrlExtractor  iris.Handler
	HeaderExtractor iris.Handler
	TokenValidator  iris.Handler
	Interceptor     iris.Handler
)

var jwtkey = []byte(config.AppConfig.GetString("app.jwtkey"))

func init() {
	TcUrlExtractor = jwt.New(tcUrlJwtConfig).Serve
	HeaderExtractor = jwt.New(headerJwtConfig).Serve

	TokenValidator = func(ctx iris.Context) {
		jwtToken, ok := ctx.Values().Get("jwt").(*jwt.Token)
		if ok {
			jwtInfo := jwtToken.Claims.(jwt.MapClaims)
			// if exp, ok := jwtInfo["exp"].(int64); ok && exp >= time.Now().Unix() {
			// 	ctx.Values().Set("user_id", jwtInfo["user_id"].(uint))
			// }
			fid := jwtInfo["user_id"].(float64)
			ctx.Values().Set("user_id", uint(fid))
		}
		ctx.Next()
	}

	Interceptor = func(ctx iris.Context) {
		if u := ctx.Values().Get("user_id"); u == nil {
			ctx.JSON(model.ErrorUnauthorized(fmt.Errorf("无法获取到token")))
		}
		ctx.Next()
	}
}

var tcUrlJwtConfig = jwt.Config{
	ErrorHandler: func(ctx iris.Context, err error) {
		if err == nil {
			return
		}
		ctx.StopExecution()
		ctx.Write(model.PermRejectedRes)
	},

	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	},

	SigningMethod: jwt.SigningMethodHS256,

	Extractor: func(ctx iris.Context) (string, error) {
		req := model.OnConnectReq{}
		ctx.ReadJSON(&req)
		tcUrl, err := url.Parse(req.TcUrl)

		return tcUrl.Query().Get("token"), err
	},
}

var headerJwtConfig = jwt.Config{
	ErrorHandler: func(ctx iris.Context, err error) {
		if err != nil {
			ctx.JSON(model.ErrorUnauthorized(err))
			ctx.StopExecution()
		}
	},

	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	},

	SigningMethod: jwt.SigningMethodHS256,
	Extractor:     jwt.FromAuthHeader,
}
