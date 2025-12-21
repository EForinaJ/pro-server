// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
	"context"

	"server/app/frontend/api/order/v1"
)

type IOrderV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	BalancePay(ctx context.Context, req *v1.BalancePayReq) (res *v1.BalancePayRes, err error)
	WechatMiniProgramPay(ctx context.Context, req *v1.WechatMiniProgramPayReq) (res *v1.WechatMiniProgramPayRes, err error)
	GetNotify(ctx context.Context, req *v1.GetNotifyReq) (res *v1.GetNotifyRes, err error)
	GetStatus(ctx context.Context, req *v1.GetStatusReq) (res *v1.GetStatusRes, err error)
}
