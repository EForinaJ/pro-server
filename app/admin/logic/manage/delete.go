package manage

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// Delete implements service.IManage.
func (s *sManage) Delete(ctx context.Context, ids []int64) (err error) {
	tx, err := g.DB().Begin(ctx)

	if err != nil {
		return utils_error.Err(response.DB_TX_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Model(dao.SysCasbin.Table()).
		Where(dao.SysCasbin.Columns().PType, "g").
		WhereIn(dao.SysCasbin.Columns().V0, ids).
		Delete()
	if err != nil {
		return err
	}

	_, err = tx.Model(dao.SysManageRole.Table()).
		Delete(dao.SysManageRole.Columns().ManageId, ids)
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}

	_, err = tx.Model(dao.SysManage.Table()).
		WhereIn(dao.SysManage.Columns().Id, ids).Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}

	return
}
