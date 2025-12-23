package distribute

import (
	"context"

	v1 "server/app/admin/api/distribute/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error) {
	err = service.Distribute().CheckCancel(ctx, req.Cancel)
	if err != nil {
		return nil, err
	}
	err = service.Distribute().Cancel(ctx, req.Cancel)
	return
}
