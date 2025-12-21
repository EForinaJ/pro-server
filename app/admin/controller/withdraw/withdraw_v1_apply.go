package withdraw

import (
	"context"

	v1 "server/app/admin/api/withdraw/v1"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

func (c *ControllerV1) Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error) {
	// 判断是否已经处理了
	ok, err := service.Withdraw().CheckStatusNotApply(ctx, req.Apply.Id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, utils_error.ErrMessage(response.FAILD, "提现已经审核")
	}

	// 判断余额是否足
	ok, err = service.Withdraw().CheckCommission(ctx, req.Apply.Id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, utils_error.ErrMessage(response.FAILD, "余额不足")
	}

	err = service.Withdraw().Apply(ctx, req.Apply)
	return
}
