package menu

import (
	"context"

	v1 "server/app/admin/api/menu/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	if req.MenuType == "menu" {
		// 判断菜单是否存在
		err = service.Menu().CheckCreate(ctx, req.Create)
		if err != nil {
			return nil, err
		}
	}

	// // 添加菜单
	err = service.Menu().Create(ctx, req.Create)
	return
}
