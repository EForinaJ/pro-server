package media

import (
	"context"

	v1 "server/app/admin/api/media/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	total, list, err := service.Media().GetList(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	res = &v1.GetListRes{
		Total: total,
		List:  list,
	}
	return
}
