package service

import (
	"context"
	dao_settlement "server/app/admin/dao/settlement"
	dto_settlement "server/app/admin/dto/settlement"
)

// 定义显示接口
type ISettlement interface {
	GetList(ctx context.Context, req *dto_settlement.Query) (total int, res []*dao_settlement.List, err error)
	GetDetail(ctx context.Context, id int64) (res *dao_settlement.Detail, err error)
}

// 定义接口变量
var localSettlement ISettlement

// 定义一个获取接口的方法
func Settlement() ISettlement {
	if localSettlement == nil {
		panic("implement not found for interface ISettlement, forgot register?")
	}
	return localSettlement
}

// 定义一个接口实现的注册方法
func RegisterSettlement(i ISettlement) {
	localSettlement = i
}
