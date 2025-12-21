// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package permission

import (
	"context"

	"server/app/admin/api/permission/v1"
)

type IPermissionV1 interface {
	GetAll(ctx context.Context, req *v1.GetAllReq) (res *v1.GetAllRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetEdit(ctx context.Context, req *v1.GetEditReq) (res *v1.GetEditRes, err error)
	Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
