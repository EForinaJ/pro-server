package onboarding

import (
	"context"

	v1 "server/app/admin/api/onboarding/v1"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

func (c *ControllerV1) Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error) {
	ok, err := service.Onboarding().CheckExistWitkey(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, utils_error.ErrMessage(response.FAILD, "该用户已入职")
	}

	err = service.Onboarding().Apply(ctx, req.Apply)
	return
}
