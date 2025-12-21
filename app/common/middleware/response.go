package middleware

import (
	"server/app/common/utils/response"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Response(r *ghttp.Request) {
	r.Middleware.Next()
	err := r.GetError()

	if err != nil {
		// 如果请求参数错误修改错误码
		if gerror.Code(err).Code() == gcode.CodeValidationFailed.Code() {

			response.Error(r).SetCode(response.PARAM_INVALID).
				SetMessage(err.Error()).Send()
		} else {
			response.Error(r).SetCode(response.ResponseCode(gerror.Code(err).Code())).
				SetMessage(err.Error()).Send()
		}

	} else {
		res := r.GetHandlerResponse()
		response.Success(r).
			SetCode(response.SUCCESS).
			SetMessage("请求成功").
			SetData(res).Send()
	}
}
