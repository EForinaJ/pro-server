package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	err = service.Account().CheckSignIn(ctx)
	if err != nil {
		return nil, err
	}

	experience, err := service.Account().SignIn(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.SignInRes{
		Experience: experience,
	}

	return
}
