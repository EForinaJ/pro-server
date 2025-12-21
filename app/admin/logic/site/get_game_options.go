package site

import (
	"context"
	dao_site "server/app/admin/dao/site"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"
)

// GetGameOptions implements service.ISite.
func (s *sSite) GetGameOptions(ctx context.Context) (res []*dao_site.Options, err error) {
	m := dao.SysGame.Ctx(ctx).
		OrderDesc(dao.SysGame.Columns().CreateTime).
		Fields(dao.SysGame.Columns().Id, dao.SysGame.Columns().Name)
	var list []*entity.SysGame
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
