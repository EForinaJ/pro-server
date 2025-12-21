package site

import (
	"context"
	dao_site "server/app/admin/dao/site"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"
)

// GetTitleOptions implements service.ISite.
func (s *sSite) GetTitleOptions(ctx context.Context, id int64) (res []*dao_site.Options, err error) {
	m := dao.SysTitle.Ctx(ctx).
		OrderDesc(dao.SysTitle.Columns().CreateTime).
		Fields(dao.SysTitle.Columns().Id, dao.SysTitle.Columns().Name)

	titleId, err := dao.SysTitle.Ctx(ctx).
		Where(dao.SysTitle.Columns().GameId, id).
		Value(dao.SysTitle.Columns().Id)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	m = m.Where(dao.SysTitle.Columns().Id, titleId)

	var list []*entity.SysTitle
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
