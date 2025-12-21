package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Recharge(ctx context.Context, req *v1.RechargeReq) (res *v1.RechargeRes, err error) {
	err = service.User().Recharge(ctx, req.Recharge)
	return
}
