package site

import (
	"context"

	v1 "server/app/admin/api/site/v1"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

func (c *ControllerV1) GetSite(ctx context.Context, req *v1.GetSiteReq) (res *v1.GetSiteRes, err error) {

	detail, err := service.Site().GetInfo(ctx)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res = &v1.GetSiteRes{
		Detail: detail,
	}
	return
}
