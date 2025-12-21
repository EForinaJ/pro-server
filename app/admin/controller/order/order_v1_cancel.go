package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error) {
	err = service.Order().CheckCancel(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = service.Order().Cancel(ctx, req.Id)
	return
}
