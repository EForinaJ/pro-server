package system

import (
	"context"

	v1 "server/app/admin/api/system/v1"
	"server/app/admin/service"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func (c *ControllerV1) GetBase(ctx context.Context, req *v1.GetBaseReq) (res *v1.GetBaseRes, err error) {

	options, err := service.System().GetOne(ctx, consts.BaseSetting)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	json, jerr := gjson.DecodeToJson(options)
	if jerr != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	jerr = json.Scan(&res)
	if jerr != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	return
}
