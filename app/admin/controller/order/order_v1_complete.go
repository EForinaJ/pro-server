package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Complete(ctx context.Context, req *v1.CompleteReq) (res *v1.CompleteRes, err error) {
	err = service.Order().CheckComplete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	err = service.Order().Complete(ctx, req.Id)
	return
}
