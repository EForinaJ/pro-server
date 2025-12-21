package title

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// Delete implements service.ITitle.
func (s *sTitle) Delete(ctx context.Context, ids []int64) (err error) {

	_, err = dao.SysTitle.Ctx(ctx).
		WhereIn(dao.SysTitle.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
