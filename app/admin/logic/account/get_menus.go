package account

import (
	"context"
	dao_account "server/app/admin/dao/account"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetMenus implements service.IAccount.
func (s *sAccount) GetMenus(ctx context.Context) (res []*dao_account.Menu, err error) {
	// 获取管理员登录的id
	userId := ctx.Value("userId")
	// 定义一个菜单实体

	// 获取管理员的角色
	roleIds, err := dao.SysManageRole.Ctx(ctx).
		Fields(dao.SysManageRole.Columns().RoleId).
		Where(dao.SysManageRole.Columns().ManageId, userId).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	// var roleIds []int64
	menuIds, err := dao.SysRoleMenu.Ctx(ctx).
		Fields(dao.SysRoleMenu.Columns().MenuId).
		WhereIn(dao.SysRoleMenu.Columns().RoleId, roleIds).
		Array()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysMenu
	err = dao.SysMenu.Ctx(ctx).
		Order(dao.SysMenu.Columns().Sort+" desc").
		Group(dao.SysMenu.Columns().Id).
		WhereIn(dao.SysMenu.Columns().Id, menuIds).Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res = make([]*dao_account.Menu, len(list))
	// 布尔值映射优化
	boolMapping := map[int]bool{
		consts.Yes: true,
		consts.Not: false,
	}
	for i, v := range list {
		var entity *dao_account.Menu
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.FAILD)
		}
		entity.IsEnable = boolMapping[v.IsEnable]
		entity.IsMenu = boolMapping[v.IsMenu]
		entity.KeepAlive = boolMapping[v.KeepAlive]
		entity.IsHide = boolMapping[v.IsHide]
		entity.IsHideTab = boolMapping[v.IsHideTab]
		entity.IsIframe = boolMapping[v.IsIframe]
		entity.ShowBadge = boolMapping[v.ShowBadge]
		entity.ShowTextBadge = boolMapping[v.ShowTextBadge]
		entity.FixedTab = boolMapping[v.FixedTab]
		entity.IsFullPage = boolMapping[v.IsFullPage]

		res[i] = entity
	}
	return
}
