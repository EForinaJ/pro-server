package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
)

func (c *ControllerV1) SavePrivacyPolicy(ctx context.Context, req *v1.SavePrivacyPolicyReq) (res *v1.SavePrivacyPolicyRes, err error) {
	err = service.System().SaveConfig(ctx, consts.PrivacyPolicy, "隐私协议", req.Content)
	return
}
