package middleware

import (
	"server/app/common/library/jwt"
	"server/app/common/utils/response"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func AdminAuth(r *ghttp.Request) {
	// 获取token
	authHeader := r.GetHeader("Authorization")

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage("访问失败,无效的token,请登录").Send()
	}
	mc, err := jwt.ParseAdminToken(parts[1])
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage("访问失败,无效的token,请登录").Send()
	}
	// 将当前请求的userID信息保存到请求的上下文c上

	r.SetCtxVar("userId", mc.Id)
	r.Middleware.Next()
}

func FrontendAuth(r *ghttp.Request) {
	// 获取token
	authHeader := r.GetHeader("Authorization")

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage("访问失败,无效的token,请登录").Send()
	}
	mc, err := jwt.ParseFrontendToken(parts[1])
	if err != nil {
		response.Error(r).SetCode(response.ACCESS_TOKEN_TIMEOUT).
			SetMessage("访问失败,无效的token,请登录").Send()
	}
	// 将当前请求的userID信息保存到请求的上下文c上

	r.SetCtxVar("userId", mc.Id)
	r.Middleware.Next()
}
