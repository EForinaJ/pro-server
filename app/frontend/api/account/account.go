// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"server/app/frontend/api/account/v1"
)

type IAccountV1 interface {
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	WxProgramBindPhone(ctx context.Context, req *v1.WxProgramBindPhoneReq) (res *v1.WxProgramBindPhoneRes, err error)
	Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error)
	ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error)
	SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error)
	GetWithdrawList(ctx context.Context, req *v1.GetWithdrawListReq) (res *v1.GetWithdrawListRes, err error)
	Withdraw(ctx context.Context, req *v1.WithdrawReq) (res *v1.WithdrawRes, err error)
	GetCardList(ctx context.Context, req *v1.GetCardListReq) (res *v1.GetCardListRes, err error)
	CreateCard(ctx context.Context, req *v1.CreateCardReq) (res *v1.CreateCardRes, err error)
	DeleteCard(ctx context.Context, req *v1.DeleteCardReq) (res *v1.DeleteCardRes, err error)
}
