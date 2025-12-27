package service

import (
	"context"
	dao_withdraw "server/app/admin/dao/withdraw"
	dto_withdraw "server/app/admin/dto/withdraw"
)

// 定义显示接口
type IWithdraw interface {
	// 检查余额是否足够
	CheckApply(ctx context.Context, req *dto_withdraw.Apply) (err error)
	GetDetail(ctx context.Context, id int64) (res *dao_withdraw.Detail, err error)
	GetList(ctx context.Context, req *dto_withdraw.Query) (total int, res []*dao_withdraw.List, err error)
	Apply(ctx context.Context, req *dto_withdraw.Apply) (err error)
}

// 定义接口变量
var localWithdraw IWithdraw

// 定义一个获取接口的方法
func Withdraw() IWithdraw {
	if localWithdraw == nil {
		panic("implement not found for interface IWithdraw, forgot register?")
	}
	return localWithdraw
}

// 定义一个接口实现的注册方法
func RegisterWithdraw(i IWithdraw) {
	localWithdraw = i
}
