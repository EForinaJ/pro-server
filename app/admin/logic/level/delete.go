package level

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// Delete implements service.ILevel.
func (s *sLevel) Delete(ctx context.Context, ids []int64) (err error) {

	_, err = dao.SysLevel.Ctx(ctx).
		WhereIn(dao.SysLevel.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
