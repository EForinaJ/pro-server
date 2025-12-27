package service

import (
	"context"
	dao_user "server/app/admin/dao/user"
	dto_user "server/app/admin/dto/user"
)

// 定义显示接口
type IUser interface {
	CheckUniquePhone(ctx context.Context, phone string) (res bool, err error)
	CheckUniquePhoneId(ctx context.Context, req *dto_user.Edit) (res bool, err error)
	CheckBalance(ctx context.Context, req *dto_user.ChangeBalance) (err error)

	GetEdit(ctx context.Context, id int64) (res *dao_user.Edit, err error)
	GetDetail(ctx context.Context, id int64) (res *dao_user.Detail, err error)
	GetList(ctx context.Context, req *dto_user.Query) (total int, res []*dao_user.List, err error)

	ChangeBalance(ctx context.Context, req *dto_user.ChangeBalance) (err error)
	Edit(ctx context.Context, req *dto_user.Edit) (err error)
	Create(ctx context.Context, req *dto_user.Create) (err error)
	Recharge(ctx context.Context, req *dto_user.Recharge) (err error)

	GetBalanceList(ctx context.Context, req *dto_user.BalanceQuery) (total int, res []*dao_user.BalanceList, err error)

	Delete(ctx context.Context, ids []int64) (err error)
}

// 定义接口变量
var localUser IUser

// 定义一个获取接口的方法
func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

// 定义一个接口实现的注册方法
func RegisterUser(i IUser) {
	localUser = i
}
