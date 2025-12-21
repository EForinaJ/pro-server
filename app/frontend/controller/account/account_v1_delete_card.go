package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) DeleteCard(ctx context.Context, req *v1.DeleteCardReq) (res *v1.DeleteCardRes, err error) {
	err = service.Account().DeleteCard(ctx, req.Id)
	return
}
