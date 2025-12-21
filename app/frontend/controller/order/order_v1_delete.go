package order

import (
	"context"

	v1 "server/app/frontend/api/order/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Order().CheckDelete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	err = service.Order().Delete(ctx, req.Id)

	return
}
