package service

import (
	"context"
	dao_account "server/app/frontend/dao/account"
	dto_account "server/app/frontend/dto/account"
)

// 定义显示接口
type IAccount interface {
	GetWorkspace(ctx context.Context) (res *dao_account.Workspace, err error)
	GetDetail(ctx context.Context) (res *dao_account.Detail, err error)

	GetCardList(ctx context.Context) (res []*dao_account.CardList, err error)
	GetWithdrawList(ctx context.Context, req *dto_account.WithdrawQuery) (total int, res []*dao_account.WithdrawList, err error)

	CheckPass(ctx context.Context, req *dto_account.Password) (err error)

	CheckSignIn(ctx context.Context) (err error)
	CheckCardUnique(ctx context.Context, req *dto_account.CreateCard) (err error)
	CheckWithdrawMoneyLessthanBalance(ctx context.Context, req *dto_account.Withdraw) (err error)

	Edit(ctx context.Context, req *dto_account.Edit) (err error)
	ChangePass(ctx context.Context, req *dto_account.Password) (err error)
	SignIn(ctx context.Context) (res int, err error)
	WxProgramBindPhone(ctx context.Context, req *dto_account.WxProgramBindPhone) (res string, err error)

	DeleteCard(ctx context.Context, id int64) (err error)
	CreateCard(ctx context.Context, req *dto_account.CreateCard) (err error)
	Withdraw(ctx context.Context, req *dto_account.Withdraw) (err error)
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
