package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) ChangeBalance(ctx context.Context, req *v1.ChangeBalanceReq) (res *v1.ChangeBalanceRes, err error) {
	if req.Mode == consts.FundLogModeSub {
		err = service.User().CheckBalance(ctx, req.ChangeBalance)
		if err != nil {
			return nil, err
		}
	}
	err = service.User().ChangeBalance(ctx, req.ChangeBalance)
	return
}
