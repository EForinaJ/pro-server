package project

import (
	"context"

	v1 "server/app/admin/api/project/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error) {
	// 判断是否可以取消
	err = service.Project().CheckCancel(ctx, req.Cancel)
	if err != nil {
		return nil, err
	}

	err = service.Project().Cancel(ctx, req.Cancel)
	return
}
