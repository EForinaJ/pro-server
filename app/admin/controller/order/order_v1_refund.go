package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Refund(ctx context.Context, req *v1.RefundReq) (res *v1.RefundRes, err error) {
	// 判断订单状态是否可以派单
	err = service.Order().CheckRefund(ctx, req.Refund)
	if err != nil {
		return nil, err
	}

	err = service.Order().Refund(ctx, req.Refund)
	return
}
