package service

import (
	"context"
	dao_auth "server/app/frontend/dao/auth"
	dto_auth "server/app/frontend/dto/auth"
)

// 定义显示接口
type IAuth interface {
	CheckLock(ctx context.Context, req *dto_auth.Login) (err error)

	// 登录
	Login(ctx context.Context, req *dto_auth.Login) (res *dao_auth.LoginRes, err error)

	// 小程序登录接口
	ProgramLogin(ctx context.Context, req *dto_auth.ProgramLogin) (res *dao_auth.LoginRes, err error)
}

// 定义接口变量
var localAuth IAuth

// 定义一个获取接口的方法
func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

// 定义一个接口实现的注册方法
func RegisterAuth(i IAuth) {
	localAuth = i
}
