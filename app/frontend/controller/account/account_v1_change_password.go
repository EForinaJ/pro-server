package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error) {
	err = service.Account().CheckPass(ctx, req.Password)

	if err != nil {
		return nil, err
	}
	err = service.Account().ChangePass(ctx, req.Password)
	return
}
