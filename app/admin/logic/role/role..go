package role

import (
	"context"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
)

type sRole struct{}

// AddRelated implements service.IRole.
func (s *sRole) AddRelated(ctx context.Context, tx gdb.TX, id int64, roleIds []int64) (err error) {
	_, err = tx.Model(dao.SysManageRole.Table()).
		Where(dao.SysManageRole.Columns().ManageId, id).Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	entites := make([]*do.SysManageRole, len(roleIds))
	for i, v := range roleIds {
		entites[i] = &do.SysManageRole{
			RoleId:   v,
			ManageId: id,
		}
	}
	_, err = tx.Model(dao.SysManageRole.Table()).
		Data(&entites).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}

func init() {
	service.RegisterRole(&sRole{})
}
