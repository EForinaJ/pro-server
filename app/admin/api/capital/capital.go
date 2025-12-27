// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package capital

import (
	"context"

	"server/app/admin/api/capital/v1"
)

type ICapitalV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
