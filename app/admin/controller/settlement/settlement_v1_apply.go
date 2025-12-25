package settlement

import (
	"context"

	v1 "server/app/admin/api/settlement/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error) {
	err = service.Settlement().CheckApply(ctx, req.Apply)
	if err != nil {
		return nil, err
	}
	err = service.Settlement().Apply(ctx, req.Apply)
	return
}
