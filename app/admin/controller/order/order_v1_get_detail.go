package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error) {
	detail, err := service.Order().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetDetailRes{
		Detail: detail,
	}

	return
}
