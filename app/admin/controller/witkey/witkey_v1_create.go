package witkey

import (
	"context"

	v1 "server/app/admin/api/witkey/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Witkey().CheckExist(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	err = service.Witkey().Create(ctx, req.Create)
	return
}
