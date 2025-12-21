package permission

import (
	"context"
	dao_permission "server/app/admin/dao/permission"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetAll implements service.IPermission.
func (s *sPermission) GetAll(ctx context.Context) (res []*dao_permission.List, err error) {
	err = dao.SysPermission.Ctx(ctx).OrderDesc(dao.SysPermission.Columns().CreateTime).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
