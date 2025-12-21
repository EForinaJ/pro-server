package game

import (
	"context"
	dao_game "server/app/admin/dao/game"
	dto_game "server/app/admin/dto/game"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

type sGame struct{}

// Delete implements service.IGame.
func (s *sGame) Delete(ctx context.Context, ids []int64) (err error) {

	_, err = dao.SysGame.Ctx(ctx).
		WhereIn(dao.SysGame.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}

// GetEdit implements service.IGame.
func (s *sGame) GetEdit(ctx context.Context, id int64) (res *dao_game.Edit, err error) {
	err = dao.SysGame.Ctx(ctx).Where(dao.SysGame.Columns().Id, id).Scan(&res)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}

// GetList implements service.IGame.
func (s *sGame) GetList(ctx context.Context, req *dto_game.Query) (total int, res []*dao_game.List, err error) {
	m := dao.SysGame.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysGame.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysGame.Columns().Name, req.Name)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysGame
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_game.List, len(list))
	for i, v := range list {
		var entity *dao_game.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		// 获取威客人数
		res[i] = entity
	}

	return
}

func init() {
	service.RegisterGame(&sGame{})
}
