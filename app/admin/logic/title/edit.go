package title

import (
	"context"
	dto_title "server/app/admin/dto/title"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.ITitle.
func (s *sTitle) Edit(ctx context.Context, req *dto_title.Edit) (err error) {

	var entity *do.SysTitle
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.UNKNOWN)
	}
	entity.Name = req.Name
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
