package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) SaveWithdraw(ctx context.Context, req *v1.SaveWithdrawReq) (res *v1.SaveWithdrawRes, err error) {
	err = service.System().SaveConfig(ctx, consts.WithdrawSetting, "提现设置", req.Withdraw)
	return
}
