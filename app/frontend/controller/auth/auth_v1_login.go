package auth

import (
	"context"

	v1 "server/app/frontend/api/auth/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {

	err = service.Auth().CheckLock(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	detail, err := service.Auth().Login(ctx, req.Login)
	if err != nil {
		return nil, err
	}
	res = &v1.LoginRes{
		LoginRes: detail,
	}
	return
}
