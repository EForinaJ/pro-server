package level

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_level "server/app/frontend/dao/level"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetAll implements service.ILevel.
func (s *sLevel) GetAll(ctx context.Context) (res []*dao_level.List, err error) {
	var list []*entity.SysLevel
	err = dao.SysLevel.Ctx(ctx).Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_level.List, len(list))
	for i, v := range list {
		var entity *dao_level.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}
