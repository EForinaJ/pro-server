package game

import (
	"context"

	v1 "server/app/admin/api/game/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	err = service.Game().CheckEdit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}

	err = service.Game().Edit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}
	return
}
