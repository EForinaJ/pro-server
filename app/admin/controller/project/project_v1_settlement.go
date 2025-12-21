package project

import (
	"context"

	v1 "server/app/admin/api/project/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Settlement(ctx context.Context, req *v1.SettlementReq) (res *v1.SettlementRes, err error) {
	// 判断是否可以结算
	err = service.Project().CheckSettlement(ctx, req.Settlement)
	if err != nil {
		return nil, err
	}

	err = service.Project().Settlement(ctx, req.Settlement)

	return
}
