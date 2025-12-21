package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) SaveBase(ctx context.Context, req *v1.SaveBaseReq) (res *v1.SaveBaseRes, err error) {
	err = service.System().SaveConfig(ctx, consts.BaseSetting, "基本设置", req.Base)
	return
}
