package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) SaveUser(ctx context.Context, req *v1.SaveUserReq) (res *v1.SaveUserRes, err error) {

	err = service.System().SaveConfig(ctx, consts.UserSetting, "用户设置", req.User)
	return
}
