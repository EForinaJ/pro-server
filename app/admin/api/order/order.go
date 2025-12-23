// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
	"context"

	"server/app/admin/api/order/v1"
)

type IOrderV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Refund(ctx context.Context, req *v1.RefundReq) (res *v1.RefundRes, err error)
	AddDiscount(ctx context.Context, req *v1.AddDiscountReq) (res *v1.AddDiscountRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Paid(ctx context.Context, req *v1.PaidReq) (res *v1.PaidRes, err error)
	Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error)
}
