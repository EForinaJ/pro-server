package site

import (
	"context"

	v1 "server/app/admin/api/site/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetTitleOptions(ctx context.Context, req *v1.GetTitleOptionsReq) (res *v1.GetTitleOptionsRes, err error) {
	list, err := service.Site().GetTitleOptions(ctx, req.GameId)
	if err != nil {
		return nil, err
	}
	res = &v1.GetTitleOptionsRes{
		List: list,
	}
	return
}
