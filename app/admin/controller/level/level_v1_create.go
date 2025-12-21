package level

import (
	"context"

	v1 "server/app/admin/api/level/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Level().CheckCreate(ctx, req.Create)
	if err != nil {
		return nil, err
	}
	err = service.Level().Create(ctx, req.Create)
	return
}
