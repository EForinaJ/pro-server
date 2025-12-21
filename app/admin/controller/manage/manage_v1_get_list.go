package manage

import (
	"context"

	v1 "server/app/admin/api/manage/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {

	total, list, err := service.Manage().GetList(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	res = &v1.GetListRes{
		List:  list,
		Total: total,
	}
	return
}
