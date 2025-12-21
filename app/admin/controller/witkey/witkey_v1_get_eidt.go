package witkey

import (
	"context"

	v1 "server/app/admin/api/witkey/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetEidt(ctx context.Context, req *v1.GetEidtReq) (res *v1.GetEidtRes, err error) {
	detail, err := service.Witkey().GetEdit(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetEidtRes{
		Edit: detail,
	}
	return
}
