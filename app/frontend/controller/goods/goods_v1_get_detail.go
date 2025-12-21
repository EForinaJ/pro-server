package goods

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"server/app/frontend/api/goods/v1"
)

func (c *ControllerV1) GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
