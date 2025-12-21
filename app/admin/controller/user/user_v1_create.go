package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	if req.Create.Phone != "" {
		ok, err := service.User().CheckUniquePhone(ctx, req.Create.Phone)
		if err != nil {
			return nil, err
		}
		if ok {
			return nil, utils_error.ErrMessage(response.FAILD, "手机号码已存在")
		}
	}

	err = service.User().Create(ctx, req.Create)
	if err != nil {
		return nil, err
	}
	return
}
