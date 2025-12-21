package order

import (
	"context"

	v1 "server/app/frontend/api/order/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) WechatMiniProgramPay(ctx context.Context, req *v1.WechatMiniProgramPayReq) (res *v1.WechatMiniProgramPayRes, err error) {
	err = service.Order().CheckPay(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	detail, err := service.Order().WeChatMiniProgramPay(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.WechatMiniProgramPayRes{
		WechatMiniProgram: detail,
	}
	return
}
