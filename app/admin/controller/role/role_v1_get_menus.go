package role

import (
	"context"

	v1 "server/app/admin/api/role/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetMenus(ctx context.Context, req *v1.GetMenusReq) (res *v1.GetMenusRes, err error) {
	ids, err := service.Menu().GetRelatedId(ctx, req.Id)
	res = &v1.GetMenusRes{
		Menus: ids,
	}
	return
}
