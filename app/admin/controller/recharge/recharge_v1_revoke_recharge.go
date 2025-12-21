package recharge

import (
	"context"

	v1 "server/app/admin/api/recharge/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) RevokeRecharge(ctx context.Context, req *v1.RevokeRechargeReq) (res *v1.RevokeRechargeRes, err error) {
	err = service.Recharge().Revoke(ctx, req.Ids)
	return
}
