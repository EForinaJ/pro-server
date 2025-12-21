package service

import (
	"context"
	dao_recharge "server/app/frontend/dao/recharge"
	dto_recharge "server/app/frontend/dto/recharge"
)

// 定义显示接口
type IRecharge interface {
	WechatMiniProgram(ctx context.Context, req *dto_recharge.WechatMiniProgram) (res *dao_recharge.WechatMiniProgram, err error)
	Notify(ctx context.Context) (err error)
	GetStatus(ctx context.Context, id int64) (status int, err error)
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
