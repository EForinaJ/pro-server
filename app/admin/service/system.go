package service

import (
	"context"
)

// 定义显示接口
type ISystem interface {
	// 获取系统配置
	GetOne(ctx context.Context, key string) (res interface{}, err error)
	// 保存配置
	SaveConfig(ctx context.Context, key, name string, value interface{}) (err error)
}

// 定义接口变量
var localSystem ISystem

// 定义一个获取接口的方法
func System() ISystem {
	if localSystem == nil {
		panic("implement not found for interface ISystem, forgot register?")
	}
	return localSystem
}

// 定义一个接口实现的注册方法
func RegisterSystem(i ISystem) {
	localSystem = i
}
