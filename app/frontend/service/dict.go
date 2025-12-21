package service

import (
	"context"
	dao_dict "server/app/frontend/dao/dict"
)

// 定义显示接口
type IDict interface {
	GetAll(ctx context.Context, code string) (res []*dao_dict.List, err error)
}

// 定义接口变量
var localDict IDict

// 定义一个获取接口的方法
func Dict() IDict {
	if localDict == nil {
		panic("implement not found for interface IDict, forgot register?")
	}
	return localDict
}

// 定义一个接口实现的注册方法
func RegisterDict(i IDict) {
	localDict = i
}
