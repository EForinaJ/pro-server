package dict_data

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "server/app/admin/api/dict_data/v1"
	"server/app/admin/service"
	"server/app/common/utils/response"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	ok, err := service.DictData().CheckUnique(ctx, req.Create.Code, req.Create.Key)
	if err != nil {
		return nil, err
	}
	if ok {
		return nil, gerror.NewCode(gcode.New(response.UNIQUE_ERROR, "", nil), "已存在")
	}

	err = service.DictData().Create(ctx, req.Create)
	return
}
