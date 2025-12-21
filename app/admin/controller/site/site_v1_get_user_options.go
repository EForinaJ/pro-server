package site

import (
	"context"

	v1 "server/app/admin/api/site/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetUserOptions(ctx context.Context, req *v1.GetUserOptionsReq) (res *v1.GetUserOptionsRes, err error) {

	list, err := service.Site().GetUserOptions(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	res = &v1.GetUserOptionsRes{
		List: list,
	}
	return
}
