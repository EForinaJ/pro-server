package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetRechargeList(ctx context.Context, req *v1.GetRechargeListReq) (res *v1.GetRechargeListRes, err error) {
	total, list, err := service.User().GetRechargeList(ctx, req.RechargeQuery)
	if err != nil {
		return nil, err
	}
	res = &v1.GetRechargeListRes{
		List:  list,
		Total: total,
	}
	return
}
