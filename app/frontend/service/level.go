package service

import (
	"context"
	dao_level "server/app/frontend/dao/level"
)

// 定义显示接口
type ILevel interface {
	GetAll(ctx context.Context) (res []*dao_level.List, err error)
}

// 定义接口变量
var localLevel ILevel

// 定义一个获取接口的方法
func Level() ILevel {
	if localLevel == nil {
		panic("implement not found for interface ILevel, forgot register?")
	}
	return localLevel
}

// 定义一个接口实现的注册方法
func RegisterLevel(i ILevel) {
	localLevel = i
}
