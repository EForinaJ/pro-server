package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) SaveWechatPay(ctx context.Context, req *v1.SaveWechatPayReq) (res *v1.SaveWechatPayRes, err error) {
	err = service.System().SaveConfig(ctx, consts.WechatPaySetting, "微信支付设置", req.WechatPay)
	return
}
