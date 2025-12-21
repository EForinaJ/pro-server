package role

import (
	"context"
	dao_role "server/app/admin/dao/role"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetAll implements service.IRole.
func (s *sRole) GetAll(ctx context.Context) (res []*dao_role.List, err error) {
	m := dao.SysRole.Ctx(ctx).
		OrderDesc(dao.SysRole.Columns().CreateTime)

	var list []*entity.SysRole
	err = m.Scan(&list)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_role.List, len(list))
	for i, v := range list {
		var entity *dao_role.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}
