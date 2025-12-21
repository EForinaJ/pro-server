package title

import (
	"context"

	v1 "server/app/admin/api/title/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetAll(ctx context.Context, req *v1.GetAllReq) (res *v1.GetAllRes, err error) {

	list, err := service.Title().GetAll(ctx, req.GameId)

	if err != nil {
		return nil, err
	}
	res = &v1.GetAllRes{
		List: list,
	}
	return
}
