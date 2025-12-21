package utils_error

import (
	"server/app/common/utils/response"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func Err(code response.ResponseCode) error {
	return gerror.NewCode(gcode.New(int(code), "", nil), response.CodeMsg(code))
}

func ErrMessage(code response.ResponseCode, message string) error {
	return gerror.NewCode(gcode.New(int(code), "", nil), message)
}
