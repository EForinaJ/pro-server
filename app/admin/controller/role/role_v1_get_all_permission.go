package role

import (
	"context"

	v1 "server/app/admin/api/role/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetAllPermission(ctx context.Context, req *v1.GetAllPermissionReq) (res *v1.GetAllPermissionRes, err error) {
	list, err := service.Permission().GetAll(ctx)
	res = &v1.GetAllPermissionRes{
		List: list,
	}
	return
}
