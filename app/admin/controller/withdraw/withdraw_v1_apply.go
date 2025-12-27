package withdraw

import (
	"context"

	v1 "server/app/admin/api/withdraw/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error) {
	err = service.Withdraw().CheckApply(ctx, req.Apply)
	err = service.Withdraw().Apply(ctx, req.Apply)
	return
}
