package middleware

import (
	"server/app/common/library/casbin"
	"server/app/common/utils/response"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// Casbin implements service.IMiddleware.
func Casbin(r *ghttp.Request) {

	obj := gstr.Replace(r.URL.Path, "/api/v1/admin", "")
	err := casbin.Enforcer.LoadPolicy()
	if err != nil {
		response.Error(r).SetCode(response.EXCEPTION).
			SetMessage("权限获取失败").Send()
	}

	sub := r.GetCtxVar("userId").String()

	ok, err := casbin.Enforcer.Enforce(sub, obj)
	if err != nil {
		response.Error(r).SetCode(response.EXCEPTION).
			SetMessage("权限获取失败").Send()
	}

	if ok {
		r.Middleware.Next()
	} else {
		response.Error(r).SetCode(response.EXCEPTION).
			SetMessage("您还没有权限").Send()
	}
}
