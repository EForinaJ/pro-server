package manage

import (
	"context"

	v1 "server/app/admin/api/manage/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetEdit(ctx context.Context, req *v1.GetEditReq) (res *v1.GetEditRes, err error) {
	detail, err := service.Manage().GetEdit(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetEditRes{
		Edit: detail,
	}
	return
}
