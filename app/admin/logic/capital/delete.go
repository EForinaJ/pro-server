package capital

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// Delete implements service.ICapital.
func (s *sCapital) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.SysCapital.Ctx(ctx).
		WhereIn(dao.SysCapital.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
