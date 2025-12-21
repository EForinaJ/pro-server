package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) Withdraw(ctx context.Context, req *v1.WithdrawReq) (res *v1.WithdrawRes, err error) {
	err = service.Account().CheckWithdrawMoneyLessthanBalance(ctx, req.Withdraw)
	if err != nil {
		return nil, err
	}
	err = service.Account().Withdraw(ctx, req.Withdraw)
	return
}
