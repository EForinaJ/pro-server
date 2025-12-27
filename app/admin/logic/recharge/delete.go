package recharge

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// Delete implements service.IRecharge.
func (s *sRecharge) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = dao.SysRecharge.Ctx(ctx).
		WhereIn(dao.SysRecharge.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
