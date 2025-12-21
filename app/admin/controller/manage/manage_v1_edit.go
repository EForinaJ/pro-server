package manage

import (
	"context"

	v1 "server/app/admin/api/manage/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	err = service.Manage().CheckPhoneUniqueById(ctx, req.Edit)
	if err != nil {
		return nil, err
	}

	err = service.Manage().Edit(ctx, req.Edit)
	return
}
