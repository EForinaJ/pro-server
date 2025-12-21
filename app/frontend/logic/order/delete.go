package order

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// Delete implements service.IOrder.
func (s *sOrder) Delete(ctx context.Context, id int64) (err error) {

	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = tx.Model(dao.SysOrderLog.Table()).
		Where(dao.SysOrderLog.Columns().OrderId, id).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, id).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
