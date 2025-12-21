package v1

import (
	dao_auth "server/app/frontend/dao/auth"
	dto_auth "server/app/frontend/dto/auth"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"登录" summary:"登录接口"`
	*dto_auth.Login
}
type LoginRes struct {
	*dao_auth.LoginRes
}

type ProgramLoginReq struct {
	g.Meta `path:"/program/login" method:"post" tags:"登录" summary:"小程序登录接口"`
	*dto_auth.ProgramLogin
}
type ProgramLoginRes struct {
	*dao_auth.LoginRes
}
