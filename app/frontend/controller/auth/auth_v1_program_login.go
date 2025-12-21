package auth

import (
	"context"

	v1 "server/app/frontend/api/auth/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) ProgramLogin(ctx context.Context, req *v1.ProgramLoginReq) (res *v1.ProgramLoginRes, err error) {

	// 获取token
	detail, err := service.Auth().ProgramLogin(ctx, req.ProgramLogin)
	if err != nil {
		return nil, err
	}
	res = &v1.ProgramLoginRes{
		LoginRes: detail,
	}
	return
}
