package manage

import (
	"context"
	dao_manage "server/app/admin/dao/manage"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetEdit implements service.IManage.
func (s *sManage) GetEdit(ctx context.Context, id int64) (res *dao_manage.Edit, err error) {
	err = dao.SysManage.Ctx(ctx).WherePri(id).Scan(&res)
	if err != nil {
		return nil, err
	}

	// 获取管理员的角色
	roleIds, err := dao.SysManageRole.Ctx(ctx).
		Fields(dao.SysManageRole.Columns().RoleId).
		Where(dao.SysManageRole.Columns().ManageId, id).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res.Roles = gconv.Int64s(roleIds)

	return
}
