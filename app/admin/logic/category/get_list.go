package category

import (
	"context"
	dao_category "server/app/admin/dao/category"
	dto_category "server/app/admin/dto/category"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.ICategory.
func (s *sCategory) GetList(ctx context.Context, req *dto_category.Query) (total int, res []*dao_category.List, err error) {
	m := dao.SysCategory.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysCategory.Columns().CreateTime)
	if req.GameId != 0 {
		m = m.Where(dao.SysCategory.Columns().GameId, req.GameId)
	}
	if req.Name != "" {
		m = m.Where(dao.SysCategory.Columns().Name, req.Name)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysCategory
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_category.List, len(list))
	for i, v := range list {
		var entity *dao_category.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id, v.GameId).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Game = game.String()

		res[i] = entity
	}

	return
}
