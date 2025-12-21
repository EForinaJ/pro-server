package menu

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
)

// Delete implements service.IMenu.
func (s *sMenu) Delete(ctx context.Context, ids []int64) (err error) {
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

	//  删除选中的菜单
	_, err = tx.Model(dao.SysMenu.Table()).WhereIn(dao.SysMenu.Columns().Id, ids).Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}

	// 修改父id 为顶层
	_, err = tx.Model(dao.SysMenu.Table()).WhereIn(dao.SysMenu.Columns().PId, ids).Data(do.SysMenu{
		PId: 0,
	}).Update()
	if err != nil {
		return utils_error.Err(response.UPDATE_FAILED)
	}

	// 删除管理表内容
	_, err = tx.Model(dao.SysRoleMenu.Table()).
		Delete(dao.SysRoleMenu.Columns().MenuId, ids)
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}
	return
}
