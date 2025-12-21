package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) CreateCard(ctx context.Context, req *v1.CreateCardReq) (res *v1.CreateCardRes, err error) {
	err = service.Account().CheckCardUnique(ctx, req.CreateCard)
	if err != nil {
		return nil, err
	}
	err = service.Account().CreateCard(ctx, req.CreateCard)
	return
}
