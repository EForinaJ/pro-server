package site

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_site "server/app/frontend/dao/site"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDictData implements service.ISite.
func (s *sSite) GetDictData(ctx context.Context, code string) (res []*dao_site.DictData, err error) {
	var list []*entity.SysDictData
	err = dao.SysDictData.Ctx(ctx).
		Where(dao.SysDictData.Columns().Code, code).Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_site.DictData, len(list))
	for i, v := range list {
		var entity *dao_site.DictData
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}
	return
}
