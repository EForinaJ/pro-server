package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) WxProgramBindPhone(ctx context.Context, req *v1.WxProgramBindPhoneReq) (res *v1.WxProgramBindPhoneRes, err error) {
	phone, err := service.Account().WxProgramBindPhone(ctx, req.WxProgramBindPhone)
	if err != nil {
		return nil, err
	}
	res = &v1.WxProgramBindPhoneRes{
		Phone: phone,
	}
	return
}
