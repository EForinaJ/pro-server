package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) SaveEmail(ctx context.Context, req *v1.SaveEmailReq) (res *v1.SaveEmailRes, err error) {
	err = service.System().SaveConfig(ctx, consts.EmailSetting, "邮箱设置", req.Email)
	return
}
