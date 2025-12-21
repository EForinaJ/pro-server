package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetAboutUs(ctx context.Context, req *v1.GetAboutUsReq) (res *v1.GetAboutUsRes, err error) {
	options, err := service.System().GetOne(ctx, consts.AboutUs)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res = &v1.GetAboutUsRes{
		Content: gconv.String(options),
	}
	return
}
