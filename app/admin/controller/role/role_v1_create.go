package role

import (
	"context"

	v1 "server/app/admin/api/role/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	//  是否存在
	err = service.Role().CheckCreate(ctx, req.Create)
	if err != nil {
		return nil, err
	}

	err = service.Role().Create(ctx, req.Create)
	if err != nil {
		return nil, err
	}

	return
}
