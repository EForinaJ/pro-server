package dict_data

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "server/app/admin/api/dict_data/v1"
	"server/app/admin/service"
	"server/app/common/utils/response"
)

func (c *ControllerV1) Edit(ctx context.Context, req *v1.EditReq) (res *v1.EditRes, err error) {
	exist, err := service.DictData().CheckUniqueById(ctx, req.Edit.Code, req.Edit.Key, req.Edit.Id)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, gerror.NewCode(gcode.New(response.UNIQUE_ERROR, "", nil), "已存在")
	}

	err = service.DictData().Edit(ctx, req.Edit)
	if err != nil {
		return nil, err
	}
	return
}
