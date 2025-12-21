package recharge

import (
	"context"

	v1 "server/app/frontend/api/recharge/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) GetStatus(ctx context.Context, req *v1.GetStatusReq) (res *v1.GetStatusRes, err error) {
	status, err := service.Recharge().GetStatus(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetStatusRes{
		Status: status,
	}
	return
}
