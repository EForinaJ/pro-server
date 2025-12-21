package role

import (
	"context"

	v1 "server/app/admin/api/role/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) EditPermissions(ctx context.Context, req *v1.EditPermissionsReq) (res *v1.EditPermissionsRes, err error) {
	err = service.Permission().AddRelated(ctx, req.Id, req.Permissions)
	return
}
