package site

import (
	"context"

	v1 "server/app/admin/api/site/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetGameOptions(ctx context.Context, req *v1.GetGameOptionsReq) (res *v1.GetGameOptionsRes, err error) {
	list, err := service.Site().GetGameOptions(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetGameOptionsRes{
		List: list,
	}
	return
}
