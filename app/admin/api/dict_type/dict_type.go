// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dict_type

import (
	"context"

	"server/app/admin/api/dict_type/v1"
)

type IDictTypeV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetEdit(ctx context.Context, req *v1.GetEditReq) (res *v1.GetEditRes, err error)
	Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
