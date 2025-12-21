package manage

import (
	"context"

	v1 "server/app/admin/api/manage/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	err = service.Manage().CheckPhoneUnique(ctx, req.Create)
	if err != nil {
		return nil, err
	}

	err = service.Manage().Create(ctx, req.Create)
	return
}
