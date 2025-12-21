package service

import (
	"context"
	dao_witkey "server/app/admin/dao/witkey"
	dto_witkey "server/app/admin/dto/witkey"
)

// 定义显示接口
type IWitkey interface {
	CheckExist(ctx context.Context, witkeyId int64) (err error)
	CheckCommission(ctx context.Context, req *dto_witkey.Commission) (err error)

	ChangeCommission(ctx context.Context, req *dto_witkey.Commission) (err error)
	Create(ctx context.Context, req *dto_witkey.Create) (err error)
	Edit(ctx context.Context, req *dto_witkey.Edit) (err error)
	Delete(ctx context.Context, ids []int64) (err error)

	GetCommissionList(ctx context.Context, req *dto_witkey.CommissionQuery) (total int, res []*dao_witkey.CommissionList, err error)
	GetDetail(ctx context.Context, id int64) (res *dao_witkey.Detail, err error)
	GetEdit(ctx context.Context, id int64) (res *dao_witkey.Edit, err error)
	GetList(ctx context.Context, req *dto_witkey.Query) (total int, res []*dao_witkey.List, err error)
}

// 定义接口变量
var localWitkey IWitkey

// 定义一个获取接口的方法
func Witkey() IWitkey {
	if localWitkey == nil {
		panic("implement not found for interface IWitkey, forgot register?")
	}
	return localWitkey
}

// 定义一个接口实现的注册方法
func RegisterWitkey(i IWitkey) {
	localWitkey = i
}
