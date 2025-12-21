package menu

import (
	"context"

	v1 "server/app/admin/api/menu/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	// 判断菜单是否存在
	err = service.Menu().CheckEdit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}

	err = service.Menu().Edit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}
	return
}
