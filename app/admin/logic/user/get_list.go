package user

import (
	"context"
	dao_user "server/app/admin/dao/user"
	dto_user "server/app/admin/dto/user"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IUser.
func (s *sUser) GetList(ctx context.Context, req *dto_user.Query) (total int, res []*dao_user.List, err error) {
	m := dao.SysUser.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysUser.Columns().CreateTime)
	if req.Id != 0 {
		m = m.WhereIn(dao.SysUser.Columns().Id, req.Id)
	}
	if req.Name != "" {
		m = m.WhereLike(dao.SysUser.Columns().Name, "%"+req.Name+"%")
	}
	if req.Phone != "" {
		m = m.Where(dao.SysUser.Columns().Phone, req.Phone)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysOrder.Columns().Status, req.Status)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysUser
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_user.List, len(list))
	for i, v := range list {
		var entity *dao_user.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		res[i] = entity
	}

	return
}
