package service

import (
	"context"
	dto_auth "server/app/admin/dto/auth"
)

// 定义显示接口
type IAuth interface {
	// 登录接口
	Login(ctx context.Context, req *dto_auth.Login) (res interface{}, err error)
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
