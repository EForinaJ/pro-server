package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) SaveContact(ctx context.Context, req *v1.SaveContactReq) (res *v1.SaveContactRes, err error) {
	err = service.System().SaveConfig(ctx, consts.ContactSetting, "客服设置", req.Contact)
	return
}
