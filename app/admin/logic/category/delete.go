package category

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// Delete implements service.ICategory.
func (s *sCategory) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.SysCategory.Ctx(ctx).
		WhereIn(dao.SysCategory.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
