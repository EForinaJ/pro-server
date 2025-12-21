package site

import (
	"context"

	v1 "server/app/frontend/api/site/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) GetSite(ctx context.Context, req *v1.GetSiteReq) (res *v1.GetSiteRes, err error) {
	detail, err := service.Site().GetInfo(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetSiteRes{
		Detail: detail,
	}
	return
}
