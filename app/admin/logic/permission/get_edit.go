package permission

import (
	"context"
	dao_permission "server/app/admin/dao/permission"
	"server/internal/dao"
)

// GetEdit implements service.IPermission.
func (s *sPermission) GetEdit(ctx context.Context, id int64) (res *dao_permission.Edit, err error) {
	err = dao.SysPermission.Ctx(ctx).WherePri(id).Scan(&res)
	return
}
