package level

import (
	"context"
	dto_level "server/app/admin/dto/level"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.ILevel.
func (s *sLevel) Edit(ctx context.Context, req *dto_level.Edit) (err error) {

	var entity *do.SysLevel
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.UNKNOWN)
	}
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysLevel.Ctx(ctx).
		Where(dao.SysLevel.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
