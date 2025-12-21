package onboarding

import (
	"context"
	dao_onboarding "server/app/admin/dao/onboarding"
	dto_onboarding "server/app/admin/dto/onboarding"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IOnboarding.
func (s *sOnboarding) GetList(ctx context.Context, req *dto_onboarding.Query) (total int, res []*dao_onboarding.List, err error) {
	m := dao.SysOnboarding.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysOnboarding.Columns().CreateTime)
	if req.Name != "" {
		userId, err := dao.SysUser.Ctx(ctx).
			WhereLike(dao.SysUser.Columns().Name, "%"+req.Name+"%").Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysOnboarding.Columns().UserId, userId)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysOnboarding.Columns().Status, req.Status)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysOnboarding
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_onboarding.List, len(list))
	for i, v := range list {
		var entity *dao_onboarding.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		game, err := dao.SysGame.Ctx(ctx).WherePri(v.GameId).Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Game = game.String()

		title, err := dao.SysTitle.Ctx(ctx).WherePri(v.TitleId).Value(dao.SysTitle.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Title = title.String()

		user, err := dao.SysUser.Ctx(ctx).WherePri(v.UserId).Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.User = user.String()

		res[i] = entity
	}
	return
}
