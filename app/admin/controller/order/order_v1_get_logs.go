package order

import (
	"context"

	v1 "server/app/admin/api/order/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetLogs(ctx context.Context, req *v1.GetLogsReq) (res *v1.GetLogsRes, err error) {
	total, list, err := service.Order().GetLogs(ctx, req.Log)
	if err != nil {
		return nil, err
	}
	res = &v1.GetLogsRes{
		List:  list,
		Total: total,
	}

	return
}
