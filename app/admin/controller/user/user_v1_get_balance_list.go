package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetBalanceList(ctx context.Context, req *v1.GetBalanceListReq) (res *v1.GetBalanceListRes, err error) {
	total, list, err := service.User().GetBalanceList(ctx, req.BalanceQuery)
	if err != nil {
		return nil, err
	}
	res = &v1.GetBalanceListRes{
		Total: total,
		List:  list,
	}

	return
}
