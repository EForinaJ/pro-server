package role

import (
	"context"
	dao_role "server/app/admin/dao/role"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetEdit implements service.IRole.
func (s *sRole) GetEdit(ctx context.Context, id int64) (res *dao_role.Edit, err error) {
	err = dao.SysRole.Ctx(ctx).WherePri(id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
