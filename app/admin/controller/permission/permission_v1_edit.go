package permission

import (
	"context"

	v1 "server/app/admin/api/permission/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	// 判断菜单是否存在
	err = service.Permission().CheckEdit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}

	// 修改菜单
	err = service.Permission().Edit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}
	return
}
