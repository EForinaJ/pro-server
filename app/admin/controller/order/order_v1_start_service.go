package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) StartService(ctx context.Context, req *v1.StartServiceReq) (res *v1.StartServiceRes, err error) {
	err = service.Order().CheckStartService(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = service.Order().StartService(ctx, req.Id)
	return
}
