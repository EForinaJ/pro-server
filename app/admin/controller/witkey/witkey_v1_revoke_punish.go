package witkey

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"server/app/admin/api/witkey/v1"
)

func (c *ControllerV1) RevokePunish(ctx context.Context, req *v1.RevokePunishReq) (res *v1.RevokePunishRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
