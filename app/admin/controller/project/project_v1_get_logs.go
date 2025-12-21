package project

import (
	"context"

	v1 "server/app/admin/api/project/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetLogs(ctx context.Context, req *v1.GetLogsReq) (res *v1.GetLogsRes, err error) {
	total, list, err := service.Project().GetLogs(ctx, req.Log)
	if err != nil {
		return nil, err
	}
	res = &v1.GetLogsRes{
		List:  list,
		Total: total,
	}
	return
}
