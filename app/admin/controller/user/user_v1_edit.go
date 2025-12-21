package user

import (
	"context"

	v1 "server/app/admin/api/user/v1"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	if req.Edit.Phone != "" {
		ok, err := service.User().CheckUniquePhoneId(ctx, req.Edit)
		if err != nil {
			return nil, err
		}
		if ok {
			return nil, utils_error.ErrMessage(response.FAILD, "手机号码已存在")
		}
	}
	err = service.User().Edit(ctx, req.Edit)
	return
}
