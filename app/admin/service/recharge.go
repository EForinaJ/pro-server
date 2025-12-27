package service

import (
	"context"
	dao_recharge "server/app/admin/dao/recharge"
	dto_recharge "server/app/admin/dto/recharge"
)

// 定义显示接口
type IRecharge interface {
	GetList(ctx context.Context, req *dto_recharge.Query) (total int, res []*dao_recharge.List, err error)
	Delete(ctx context.Context, ids []int64) (err error)
}

// 定义接口变量
var localRecharge IRecharge

// 定义一个获取接口的方法
func Recharge() IRecharge {
	if localRecharge == nil {
		panic("implement not found for interface IRecharge, forgot register?")
	}
	return localRecharge
}

// 定义一个接口实现的注册方法
func RegisterRecharge(i IRecharge) {
	localRecharge = i
}
