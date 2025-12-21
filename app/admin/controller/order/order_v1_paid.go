package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Paid(ctx context.Context, req *v1.PaidReq) (res *v1.PaidRes, err error) {
	err = service.Order().CheckPaid(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = service.Order().Paid(ctx, req.Id)
	return
}
