package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/service"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

func (c *ControllerV1) ChangeBalance(ctx context.Context, req *v1.ChangeBalanceReq) (res *v1.ChangeBalanceRes, err error) {
	if req.Mode == consts.FundLogModeSub {
		ok, err := service.User().CheckBalance(ctx, req.ChangeBalance.Id, req.ChangeBalance.Money)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, utils_error.ErrMessage(response.FAILD, "余额不足，减少金额超出余额")
		}
	}
	err = service.User().ChangeBalance(ctx, req.ChangeBalance)
	return
}
