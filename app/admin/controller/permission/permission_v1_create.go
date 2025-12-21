package permission

import (
	"context"

	v1 "server/app/admin/api/permission/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 判断菜单是否存在
	err = service.Permission().CheckCreate(ctx, req.Create)
	if err != nil {
		return nil, err
	}

	// 添加菜单
	err = service.Permission().Create(ctx, req.Create)
	if err != nil {
		return nil, err
	}
	return
}
