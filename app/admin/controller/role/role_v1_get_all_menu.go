package role

import (
	"context"

	v1 "server/app/admin/api/role/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetAllMenu(ctx context.Context, req *v1.GetAllMenuReq) (res *v1.GetAllMenuRes, err error) {
	list, err := service.Menu().GetAll(ctx)
	res = &v1.GetAllMenuRes{
		List: list,
	}
	return
}
