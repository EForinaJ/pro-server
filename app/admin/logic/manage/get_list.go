package manage

import (
	"context"
	dao_manage "server/app/admin/dao/manage"
	dto_manage "server/app/admin/dto/manage"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IManage.
func (s *sManage) GetList(ctx context.Context, req *dto_manage.Query) (total int, res []*dao_manage.List, err error) {
	m := dao.SysManage.Ctx(ctx).OrderDesc(dao.SysManage.Columns().CreateTime)
	if req.Name != "" {
		m = m.WhereLike(dao.SysManage.Columns().Name, req.Name)
	}
	m = m.Page(req.Page, req.Limit)
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	err = m.Scan(&res)

	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	for i, v := range res {
		rs, err := dao.SysManageRole.Ctx(ctx).
			WhereIn(dao.SysManageRole.Columns().ManageId, v.Id).
			Fields(dao.SysManageRole.Columns().RoleId).
			Array()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		titles, err := dao.SysRole.Ctx(ctx).
			WhereIn(dao.SysRole.Columns().Id, rs).
			Fields(dao.SysRole.Columns().Name).
			Array()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i].Roles = gconv.Strings(titles)
	}

	return
}
