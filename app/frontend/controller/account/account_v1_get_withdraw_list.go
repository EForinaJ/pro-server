package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) GetWithdrawList(ctx context.Context, req *v1.GetWithdrawListReq) (res *v1.GetWithdrawListRes, err error) {
	total, list, err := service.Account().GetWithdrawList(ctx, req.WithdrawQuery)
	if err != nil {
		return nil, err
	}
	res = &v1.GetWithdrawListRes{
		Total: total,
		List:  list,
	}
	return
}
