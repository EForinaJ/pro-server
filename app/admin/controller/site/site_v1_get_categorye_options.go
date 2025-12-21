package site

import (
	"context"

	v1 "server/app/admin/api/site/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetCategoryeOptions(ctx context.Context, req *v1.GetCategoryeOptionsReq) (res *v1.GetCategoryeOptionsRes, err error) {
	list, err := service.Site().GetCategoryOptions(ctx, req.GameId)
	if err != nil {
		return nil, err
	}
	res = &v1.GetCategoryeOptionsRes{
		List: list,
	}
	return
}
