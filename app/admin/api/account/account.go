// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	"server/app/admin/api/account/v1"
)

type IAccountV1 interface {
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	GetMenu(ctx context.Context, req *v1.GetMenuReq) (res *v1.GetMenuRes, err error)
	GetDashboard(ctx context.Context, req *v1.GetDashboardReq) (res *v1.GetDashboardRes, err error)
}
