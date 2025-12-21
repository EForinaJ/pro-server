package site

import (
	"context"
	dao_site "server/app/admin/dao/site"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"
)

// GetUserOptions implements service.ISite.
func (s *sSite) GetUserOptions(ctx context.Context, name string) (res []*dao_site.Options, err error) {
	m := dao.SysUser.Ctx(ctx).
		Page(1, 5).
		OrderDesc(dao.SysUser.Columns().CreateTime).
		Fields(dao.SysUser.Columns().Id, dao.SysUser.Columns().Name)

	userId, err := dao.SysUser.Ctx(ctx).
		WhereLike(dao.SysUser.Columns().Name, "%"+name+"%").
		Value(dao.SysUser.Columns().Id)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	m = m.Where(dao.SysUser.Columns().Id, userId)

	var list []*entity.SysUser
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
