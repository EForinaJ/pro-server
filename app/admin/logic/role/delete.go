package role

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// Delete implements service.IRole.
func (s *sRole) Delete(ctx context.Context, ids []int64) (err error) {
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

	roles, err := tx.Model(dao.SysRole.Table()).
		Fields(dao.SysRole.Columns().Code).
		WhereIn(dao.SysRole.Columns().Id, ids).Array()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}

	_, err = tx.Model(dao.SysCasbin.Table()).
		WhereIn(dao.SysCasbin.Columns().V1, roles).
		Delete()
	if err != nil {
		return err
	}

	_, err = tx.Model(dao.SysRoleMenu.Table()).
		Delete(dao.SysRoleMenu.Columns().RoleId, ids)
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}

	_, err = tx.Model(dao.SysRolePermission.Table()).
		Delete(dao.SysRolePermission.Columns().RoleId, ids)
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}

	_, err = tx.Model(dao.SysRole.Table()).
		WhereIn(dao.SysRole.Columns().Id, ids).Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}

	return
}
