package service

import (
	"context"
	dao_aftersales "server/app/admin/dao/aftersales"
	dto_aftersales "server/app/admin/dto/aftersales"
)

// 定义显示接口
type IAftersales interface {
	GetList(ctx context.Context, req *dto_aftersales.Query) (total int, res []*dao_aftersales.List, err error)
	GetDetail(ctx context.Context, id int64) (res *dao_aftersales.Detail, err error)
	// Apply(ctx context.Context, req *dto_aftersales.Apply) (err error)
	// CheckApply(ctx context.Context, req *dto_aftersales.Apply) (err error)
}

// 定义接口变量
var localAftersales IAftersales

// 定义一个获取接口的方法
func Aftersales() IAftersales {
	if localAftersales == nil {
		panic("implement not found for interface IAftersales, forgot register?")
	}
	return localAftersales
}

// 定义一个接口实现的注册方法
func RegisterAftersales(i IAftersales) {
	localAftersales = i
}
