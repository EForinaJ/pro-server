// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"server/app/admin/api/role/v1"
)

type IRoleV1 interface {
	GetAllMenu(ctx context.Context, req *v1.GetAllMenuReq) (res *v1.GetAllMenuRes, err error)
	GetAllPermission(ctx context.Context, req *v1.GetAllPermissionReq) (res *v1.GetAllPermissionRes, err error)
	EditPermissions(ctx context.Context, req *v1.EditPermissionsReq) (res *v1.EditPermissionsRes, err error)
	GetPermissions(ctx context.Context, req *v1.GetPermissionsReq) (res *v1.GetPermissionsRes, err error)
	EditMenus(ctx context.Context, req *v1.EditMenusReq) (res *v1.EditMenusRes, err error)
	GetMenus(ctx context.Context, req *v1.GetMenusReq) (res *v1.GetMenusRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	GetEdit(ctx context.Context, req *v1.GetEditReq) (res *v1.GetEditRes, err error)
	Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}
