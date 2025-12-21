package account

import (
	"context"
	dao_account "server/app/admin/dao/account"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IAccount.
func (s *sAccount) GetDetail(ctx context.Context) (res *dao_account.Detail, err error) {
	userId := ctx.Value("userId")
	options, err := g.Redis().Get(ctx, consts.AdminAccout+gconv.String(userId))
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}

	if !options.IsEmpty() {
		err = options.Scan(&res)
		if err != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR)
		}
		return
	}

	err = dao.SysManage.Ctx(ctx).Fields(dao.SysManage.Columns().Id,
		dao.SysManage.Columns().Name,
		dao.SysManage.Columns().Id,
		dao.SysManage.Columns().Avatar).
		Where(dao.SysManage.Columns().Id, userId).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	// 获取角色
	// roles, err := dao.SysManageRole.
	// 	Ctx(ctx).
	// 	Where(dao.SysManageRole.Columns().ManageId, userId).
	// 	Fields(dao.SysManageRole.Columns().RoleId).
	// 	Array()
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }

	// permissions, err := dao.SysRolePermission.Ctx(ctx).
	// 	WhereIn(dao.SysRolePermission.Columns().RoleId, roles).
	// 	Fields(dao.SysRolePermission.Columns().PermissionId).
	// 	Array()
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }

	// permissionKey, err := dao.SysPermission.Ctx(ctx).
	// 	WhereIn(dao.SysPermission.Columns().Id, permissions).
	// 	Fields(dao.SysPermission.Columns().Permission).
	// 	Array()
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }
	// res.Permission = gconv.Strings(permissionKey)

	err = g.Redis().SetEX(ctx, consts.AdminAccout+gconv.String(userId), res, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR)
	}

	return
}
