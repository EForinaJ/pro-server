package account

import (
	"context"

	v1 "server/app/admin/api/account/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetMenu(ctx context.Context, req *v1.GetMenuReq) (res *v1.GetMenuRes, err error) {
	list, err := service.Account().GetMenus(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetMenuRes{
		List: list,
	}
	return
}
