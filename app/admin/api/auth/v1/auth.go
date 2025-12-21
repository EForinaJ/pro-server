package v1

import (
	dto_auth "server/app/admin/dto/auth"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"登录" summary:"登录接口"`
	*dto_auth.Login
}
type LoginRes struct {
	Token string `json:"token" dc:"登录Token"`
}
