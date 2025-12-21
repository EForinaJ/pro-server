package site

import (
	"context"
	dao_site "server/app/admin/dao/site"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"
)

// GetCategoryOptions implements service.ISite.
func (s *sSite) GetCategoryOptions(ctx context.Context, id int64) (res []*dao_site.Options, err error) {
	m := dao.SysCategory.Ctx(ctx).
		OrderDesc(dao.SysCategory.Columns().CreateTime).
		Fields(dao.SysCategory.Columns().Id, dao.SysCategory.Columns().Name)

	categoryId, err := dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().GameId, id).
		Value(dao.SysCategory.Columns().Id)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	m = m.Where(dao.SysCategory.Columns().Id, categoryId)

	var list []*entity.SysCategory
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	res = make([]*dao_site.Options, len(list))
	for i, v := range list {
		res[i] = &dao_site.Options{
			Name:  v.Name,
			Id:    v.Id,
			Value: v.Id,
		}
	}
	return
}
