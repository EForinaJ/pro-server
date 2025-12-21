// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package project

import (
	"context"

	"server/app/admin/api/project/v1"
)

type IProjectV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetLogs(ctx context.Context, req *v1.GetLogsReq) (res *v1.GetLogsRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Settlement(ctx context.Context, req *v1.SettlementReq) (res *v1.SettlementRes, err error)
	Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
