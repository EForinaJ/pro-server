package order

import (
	"context"

	v1 "server/app/frontend/api/order/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	total, list, err := service.Order().GetList(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	res = &v1.GetListRes{
		Total: total,
		List:  list,
	}
	return
}
