package dict_type

import (
	"context"

	v1 "server/app/admin/api/dict_type/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.DictType().Delete(ctx, req.Ids)
	return
}
