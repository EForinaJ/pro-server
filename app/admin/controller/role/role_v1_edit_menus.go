package role

import (
	"context"

	v1 "server/app/admin/api/role/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) EditMenus(ctx context.Context, req *v1.EditMenusReq) (res *v1.EditMenusRes, err error) {
	err = service.Menu().AddRelated(ctx, req.Id, req.Menus)
	return
}
