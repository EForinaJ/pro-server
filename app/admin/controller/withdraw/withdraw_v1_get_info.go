package withdraw

import (
	"context"

	v1 "server/app/admin/api/withdraw/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	detail, err := service.Withdraw().GetInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetInfoRes{
		Detail: detail,
	}
	return
}
