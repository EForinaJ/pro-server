package level

import (
	"context"
	dao_level "server/app/admin/dao/level"
	dto_level "server/app/admin/dto/level"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.ILevel.
func (s *sLevel) GetList(ctx context.Context, req *dto_level.Query) (total int, res []*dao_level.List, err error) {
	m := dao.SysLevel.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysLevel.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysLevel.Columns().Name, req.Name)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysLevel
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_level.List, len(list))
	for i, v := range list {
		var entity *dao_level.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}
