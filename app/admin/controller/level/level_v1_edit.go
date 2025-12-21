package level

import (
	"context"

	v1 "server/app/admin/api/level/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {

	err = service.Level().CheckEdit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}
	err = service.Level().Edit(ctx, req.Edit)
	return
}
