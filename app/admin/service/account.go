package service

import (
	"context"
	dao_account "server/app/admin/dao/account"
)

// 定义显示接口
type IAccount interface {
	// 账户信息
	GetDetail(ctx context.Context) (res *dao_account.Detail, err error)
	// 账户菜单
	GetMenus(ctx context.Context) (res []*dao_account.Menu, err error)
	// 获取统计数据
	GetDashboard(ctx context.Context) (res *dao_account.Dashboard, err error)
}

// 定义接口变量
var localAccount IAccount

// 定义一个获取接口的方法
func Account() IAccount {
	if localAccount == nil {
		panic("implement not found for interface IAccount, forgot register?")
	}
	return localAccount
}

// 定义一个接口实现的注册方法
func RegisterAccount(i IAccount) {
	localAccount = i
}
