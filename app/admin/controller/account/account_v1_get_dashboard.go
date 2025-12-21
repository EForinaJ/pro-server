package account

import (
	"context"

	v1 "server/app/admin/api/account/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetDashboard(ctx context.Context, req *v1.GetDashboardReq) (res *v1.GetDashboardRes, err error) {
	detail, err := service.Account().GetDashboard(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetDashboardRes{
		Dashboard: detail,
	}
	return
}
