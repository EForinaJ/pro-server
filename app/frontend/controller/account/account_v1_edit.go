package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	err = service.Account().Edit(ctx, req.Edit)
	return
}
