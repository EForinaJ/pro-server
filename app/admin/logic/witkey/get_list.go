package witkey

import (
	"context"
	dao_witkey "server/app/admin/dao/witkey"
	dto_witkey "server/app/admin/dto/witkey"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IWitkey.
func (s *sWitkey) GetList(ctx context.Context, req *dto_witkey.Query) (total int, res []*dao_witkey.List, err error) {
	m := dao.SysWitkey.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysWitkey.Columns().CreateTime)
	if req.Id != 0 {
		m = m.WhereIn(dao.SysWitkey.Columns().Id, req.Id)
	}
	if req.Name != "" {
		user, err := dao.SysUser.Ctx(ctx).
			WhereLike(dao.SysUser.Columns().Name, "%"+req.Name+"%").Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysWitkey.Columns().UserId, user.Int64())
	}
	if req.Phone != "" {
		user, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Phone, req.Phone).
			Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysWitkey.Columns().UserId, user.Int64())
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysWitkey
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_witkey.List, len(list))
	for i, v := range list {
		var entity *dao_witkey.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, v.UserId).
			Fields(dao.SysUser.Columns().Name, dao.SysUser.Columns().Avatar, dao.SysUser.Columns().Phone).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		var userEntity *dao_witkey.User
		err = gconv.Scan(user, &userEntity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.User = userEntity

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id, v.GameId).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Game = game.String()

		title, err := dao.SysTitle.Ctx(ctx).
			Where(dao.SysTitle.Columns().Id, v.TitleId).
			Value(dao.SysTitle.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Title = title.String()

		res[i] = entity
	}

	return
}
