package project

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// Delete implements service.IProject.
func (s *sProject) Delete(ctx context.Context, ids []int64) (err error) {
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

	_, err = tx.Model(dao.SysProjectLog.Table()).
		WhereIn(dao.SysProjectLog.Columns().ProjectId, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = tx.Model(dao.SysProject.Table()).
		WhereIn(dao.SysProject.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
