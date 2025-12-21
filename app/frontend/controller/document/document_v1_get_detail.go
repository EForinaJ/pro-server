package document

import (
	"context"

	v1 "server/app/frontend/api/document/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error) {
	detail, err := service.Document().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &v1.GetDetailRes{
		Detail: detail,
	}
	return
}
