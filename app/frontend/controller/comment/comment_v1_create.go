package comment

import (
	"context"

	v1 "server/app/frontend/api/comment/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Comment().CheckCreate(ctx, req.Create)
	if err != nil {
		return nil, err
	}

	err = service.Comment().Create(ctx, req.Create)
	return
}
