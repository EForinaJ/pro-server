package category

import (
	"context"

	v1 "server/app/admin/api/category/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetEdit(ctx context.Context, req *v1.GetEditReq) (res *v1.GetEditRes, err error) {
	detail, err := service.Category().GetEdit(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetEditRes{
		Edit: detail,
	}
	return
}
