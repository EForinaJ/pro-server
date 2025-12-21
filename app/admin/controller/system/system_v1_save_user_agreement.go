package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) SaveUserAgreement(ctx context.Context, req *v1.SaveUserAgreementReq) (res *v1.SaveUserAgreementRes, err error) {
	err = service.System().SaveConfig(ctx, consts.UserAgreement, "用户协议", req.Content)
	return
}
