package service

import (
	"context"
	dao_order "server/app/frontend/dao/order"
	dto_order "server/app/frontend/dto/order"
)

// 定义显示接口
type IOrder interface {
	GetDetail(ctx context.Context, id int64) (res *dao_order.Detail, err error)
	GetList(ctx context.Context, req *dto_order.Query) (total int, res []*dao_order.List, err error)
	GetStatus(ctx context.Context, id int64) (status int, err error)

	CheckCancel(ctx context.Context, id int64) (err error)
	CheckDelete(ctx context.Context, id int64) (err error)
	CheckPay(ctx context.Context, id int64) (err error)
	CheckBalanceSufficient(ctx context.Context, id int64) (err error)

	Cancel(ctx context.Context, id int64) (err error)
	Delete(ctx context.Context, id int64) (err error)

	BalancePay(ctx context.Context, id int64) (err error)
	WeChatMiniProgramPay(ctx context.Context, id int64) (res *dao_order.WechatMiniProgram, err error)
	Notify(ctx context.Context) (err error)
}

// 定义接口变量
var localOrder IOrder

// 定义一个获取接口的方法
func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

// 定义一个接口实现的注册方法
func RegisterOrder(i IOrder) {
	localOrder = i
}
