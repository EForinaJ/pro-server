package witkey

import (
	"context"

	v1 "server/app/admin/api/witkey/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetCommissionList(ctx context.Context, req *v1.GetCommissionListReq) (res *v1.GetCommissionListRes, err error) {
	total, list, err := service.Witkey().GetCommissionList(ctx, req.CommissionQuery)
	if err != nil {
		return nil, err
	}
	res = &v1.GetCommissionListRes{
		Total: total,
		List:  list,
	}

	return
}
