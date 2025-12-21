package witkey

import (
	"context"
	dao_witkey "server/app/admin/dao/witkey"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IWitkey.
func (s *sWitkey) GetDetail(ctx context.Context, id int64) (res *dao_witkey.Detail, err error) {
	info, err := dao.SysWitkey.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if info.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}
	// 2. 使用结构体转换代替手动映射
	var detail *dao_witkey.Detail
	if err := gconv.Scan(info.Map(), &detail); err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	userModel, err := dao.SysUser.Ctx(ctx).
		Fields(dao.SysUser.Columns().Name, dao.SysUser.Columns().Avatar, dao.SysUser.Columns().Phone).
		Where(dao.SysUser.Columns().Id, info.GMap().Get(dao.SysWitkey.Columns().UserId)).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	// 2. 使用结构体转换代替手动映射
	var user *dao_witkey.User
	if err := gconv.Scan(userModel.Map(), &user); err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.User = user

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, info.GMap().Get(dao.SysWitkey.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Game = game.String()

	title, err := dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().Id, info.GMap().Get(dao.SysWitkey.Columns().TitleId)).
		Value(dao.SysTitle.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	detail.Title = title.String()

	return detail, nil
}
