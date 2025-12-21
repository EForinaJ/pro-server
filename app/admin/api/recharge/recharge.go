// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package recharge

import (
	"context"

	"server/app/admin/api/recharge/v1"
)

type IRechargeV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	RevokeRecharge(ctx context.Context, req *v1.RevokeRechargeReq) (res *v1.RevokeRechargeRes, err error)
}
