package role

import (
	"context"
	dao_role "server/app/admin/dao/role"
	dto_role "server/app/admin/dto/role"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IRole.
func (s *sRole) GetList(ctx context.Context, req *dto_role.Query) (total int, res []*dao_role.List, err error) {
	m := dao.SysRole.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysRole.Columns().CreateTime)

	if req.Name != "" {
		m = m.Where(dao.SysRole.Columns().Name, req.Name)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysRole
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_role.List, len(list))
	for i, v := range list {
		var entity *dao_role.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}
