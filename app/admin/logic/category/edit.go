package category

import (
	"context"
	dto_category "server/app/admin/dto/category"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Edit implements service.ICategory.
func (s *sCategory) Edit(ctx context.Context, req *dto_category.Edit) (err error) {

	var entity *do.SysCategory
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.UNKNOWN)
	}
	entity.UpdateTime = gtime.Now()
	_, err = dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().Id, req.Id).
		Data(&entity).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
