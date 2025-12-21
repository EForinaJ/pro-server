// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package withdraw

import (
	"context"

	"server/app/admin/api/withdraw/v1"
)

type IWithdrawV1 interface {
	GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error)
}
