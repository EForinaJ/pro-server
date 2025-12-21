package dict_type

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "server/app/admin/api/dict_type/v1"
	"server/app/admin/service"
	"server/app/common/utils/response"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	exist, err := service.DictType().CheckUniqueById(ctx, req.Edit.Code, req.Edit.Id)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, gerror.NewCode(gcode.New(response.UNIQUE_ERROR, "", nil), "已存在")
	}

	err = service.DictType().Edit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}
	return
}
