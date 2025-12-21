package title

import (
	"context"
	dao_title "server/app/admin/dao/title"
	dto_title "server/app/admin/dto/title"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.ITitle.
func (s *sTitle) GetList(ctx context.Context, req *dto_title.Query) (total int, res []*dao_title.List, err error) {
	m := dao.SysTitle.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysTitle.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysTitle.Columns().Name, req.Name)
	}

	if req.GameId != 0 {
		m = m.Where(dao.SysTitle.Columns().GameId, req.GameId)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysTitle
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_title.List, len(list))
	for i, v := range list {
		var entity *dao_title.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Name = v.Name

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id, v.GameId).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		entity.Game = gconv.String(game)
		res[i] = entity
	}

	return
}
