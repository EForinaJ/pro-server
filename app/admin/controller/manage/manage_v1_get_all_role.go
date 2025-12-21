package manage

import (
	"context"

	v1 "server/app/admin/api/manage/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetAllRole(ctx context.Context, req *v1.GetAllRoleReq) (res *v1.GetAllRoleRes, err error) {
	list, err := service.Role().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.GetAllRoleRes{
		List: list,
	}
	return
}
