package order

import (
	"context"

	v1 "server/app/frontend/api/order/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) BalancePay(ctx context.Context, req *v1.BalancePayReq) (res *v1.BalancePayRes, err error) {
	err = service.Order().CheckPay(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	err = service.Order().CheckBalanceSufficient(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = service.Order().BalancePay(ctx, req.Id)

	return
}
