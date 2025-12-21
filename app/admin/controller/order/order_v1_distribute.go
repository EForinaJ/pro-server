package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Distribute(ctx context.Context, req *v1.DistributeReq) (res *v1.DistributeRes, err error) {
	// 判断是否已有派发过
	err = service.Order().CheckDistribute(ctx, req.Distribute)
	if err != nil {
		return nil, err
	}

	err = service.Order().Distribute(ctx, req.Distribute)
	return
}
