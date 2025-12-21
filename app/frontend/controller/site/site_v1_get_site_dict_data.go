package site

import (
	"context"

	v1 "server/app/frontend/api/site/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) GetSiteDictData(ctx context.Context, req *v1.GetSiteDictDataReq) (res *v1.GetSiteDictDataRes, err error) {
	list, err := service.Site().GetDictData(ctx, req.Code)
	if err != nil {
		return nil, err
	}
	res = &v1.GetSiteDictDataRes{
		List: list,
	}
	return
}
