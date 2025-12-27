package capital

import (
	"context"

	v1 "server/app/admin/api/capital/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.Capital().Delete(ctx, req.Ids)
	return
}
