// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"server/app/admin/api/user/v1"
)

type IUserV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetEidt(ctx context.Context, req *v1.GetEidtReq) (res *v1.GetEidtRes, err error)
	Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error)
	ChangeBalance(ctx context.Context, req *v1.ChangeBalanceReq) (res *v1.ChangeBalanceRes, err error)
	Recharge(ctx context.Context, req *v1.RechargeReq) (res *v1.RechargeRes, err error)
	GetRechargeList(ctx context.Context, req *v1.GetRechargeListReq) (res *v1.GetRechargeListRes, err error)
	GetBalanceList(ctx context.Context, req *v1.GetBalanceListReq) (res *v1.GetBalanceListRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GetLogs(ctx context.Context, req *v1.GetLogsReq) (res *v1.GetLogsRes, err error)
}
