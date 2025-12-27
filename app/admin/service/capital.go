package service

import (
	"context"
	dao_capital "server/app/admin/dao/capital"
	dto_capital "server/app/admin/dto/capital"
)

// 定义显示接口
type ICapital interface {
	GetList(ctx context.Context, req *dto_capital.Query) (total int, res []*dao_capital.List, err error)
	Delete(ctx context.Context, ids []int64) (err error)
}

// 定义接口变量
var localCapital ICapital

// 定义一个获取接口的方法
func Capital() ICapital {
	if localCapital == nil {
		panic("implement not found for interface ICapital, forgot register?")
	}
	return localCapital
}

// 定义一个接口实现的注册方法
func RegisterCapital(i ICapital) {
	localCapital = i
}
