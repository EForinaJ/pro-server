package dict_data

import (
	"context"

	v1 "server/app/admin/api/dict_data/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetAll(ctx context.Context, req *v1.GetAllReq) (res *v1.GetAllRes, err error) {
	list, err := service.DictData().GetAll(ctx, req.Code)
	if err != nil {
		return nil, err
	}
	res = &v1.GetAllRes{
		List: list,
	}
	return
}
