package permission

import (
	"context"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sPermission struct{}

// GetRelatedId implements service.IPermission.
func (s *sPermission) GetRelatedId(ctx context.Context, id int64) (ids []int64, err error) {
	Ids, err := dao.SysRolePermission.Ctx(ctx).Where(dao.SysRolePermission.Columns().RoleId, id).
		Fields(dao.SysRolePermission.Columns().PermissionId).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	ids = gconv.Int64s(Ids)
	return
}

// AddRelated implements service.IPermission.
func (s *sPermission) AddRelated(ctx context.Context, id int64, permission []int64) (err error) {
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

	_, err = tx.Model(dao.SysRolePermission.Table()).
		Where(dao.SysRoleMenu.Columns().RoleId, id).Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}
	entites := make([]*do.SysRolePermission, len(permission))
	for i, v := range permission {
		entites[i] = &do.SysRolePermission{
			PermissionId: v,
			RoleId:       id,
		}
	}
	_, err = tx.Model(dao.SysRolePermission.Table()).
		Data(&entites).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	code, err := tx.Model(dao.SysRole.Table()).Where(dao.SysRole.Columns().Id, id).
		Value(dao.SysRole.Columns().Code)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	// 更新casbin权限内容
	_, err = tx.Model(dao.SysCasbin.Table()).
		Where(dao.SysCasbin.Columns().PType, "p").
		Where(dao.SysCasbin.Columns().V0, code.String()).Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}

	pers, err := tx.Model(dao.SysPermission.Table()).
		Fields(dao.SysPermission.Columns().Permission).
		WhereIn(dao.SysPermission.Columns().Id, permission).
		Array()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	pentites := make([]g.Map, len(pers))
	for i, v := range pers {
		pentites[i] = g.Map{
			dao.SysCasbin.Columns().PType: "p",
			dao.SysCasbin.Columns().V0:    code,
			dao.SysCasbin.Columns().V1:    v.String(),
		}
	}

	_, err = tx.Model(dao.SysCasbin.Table()).
		Data(pentites).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}

// Delete implements service.IPermission.
func (s *sPermission) Delete(ctx context.Context, ids []int64) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Model(dao.SysRolePermission.Table()).
		WhereIn(dao.SysRolePermission.Columns().PermissionId, ids).
		Delete()
	if err != nil {
		return err
	}

	pers, err := tx.Model(dao.SysPermission.Table()).
		Fields(dao.SysPermission.Columns().Permission).
		WhereIn(dao.SysPermission.Columns().Id, ids).Array()
	if err != nil {
		return err
	}

	_, err = tx.Model(dao.SysCasbin.Table()).
		WhereIn(dao.SysCasbin.Columns().V1, pers).
		Delete()
	if err != nil {
		return err
	}

	_, err = tx.Model(dao.SysPermission.Table()).
		WhereIn(dao.SysPermission.Columns().Pid, ids).
		Data(do.SysPermission{
			Pid: 0,
		}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED)
	}

	_, err = tx.Model(dao.SysPermission.Table()).
		WhereIn(dao.SysPermission.Columns().Id, ids).
		Delete()
	if err != nil {
		return err
	}

	return
}

func init() {
	service.RegisterPermission(&sPermission{})
}
