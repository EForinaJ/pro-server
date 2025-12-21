package title

import (
	"context"

	v1 "server/app/admin/api/title/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	err = service.Title().CheckEdit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}

	err = service.Title().Edit(ctx, req.Edit)

	return
}
