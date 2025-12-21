package onboarding

import (
	"context"
	dao_onboarding "server/app/admin/dao/onboarding"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IOnboarding.
func (s *sOnboarding) GetDetail(ctx context.Context, id int64) (res *dao_onboarding.Detail, err error) {

	info, err := dao.SysOnboarding.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if info.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}

	// 2. 使用结构体转换代替手动映射
	var detail *dao_onboarding.Detail
	if err := gconv.Scan(info.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	title, err := dao.SysTitle.Ctx(ctx).
		WherePri(info.GMap().Get(dao.SysOnboarding.Columns().TitleId)).
		Value(dao.SysTitle.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Title = title.String()

	game, err := dao.SysGame.Ctx(ctx).
		WherePri(info.GMap().Get(dao.SysOnboarding.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Game = game.String()

	user, err := dao.SysUser.Ctx(ctx).
		WherePri(info.GMap().Get(dao.SysOnboarding.Columns().UserId)).
		Value(dao.SysUser.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.User = user.String()

	manage, err := dao.SysManage.Ctx(ctx).
		WherePri(info.GMap().Get(dao.SysOnboarding.Columns().ManageId)).
		Value(dao.SysManage.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Manage = manage.String()

	return detail, nil
}
