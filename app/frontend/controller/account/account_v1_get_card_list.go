package account

import (
	"context"

	v1 "server/app/frontend/api/account/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) GetCardList(ctx context.Context, req *v1.GetCardListReq) (res *v1.GetCardListRes, err error) {
	list, err := service.Account().GetCardList(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetCardListRes{
		List: list,
	}
	return
}
