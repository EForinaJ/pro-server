package recharge

import (
	"context"
	dao_recharge "server/app/admin/dao/recharge"
	dto_recharge "server/app/admin/dto/recharge"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IRecharge.
func (s *sRecharge) GetList(ctx context.Context, req *dto_recharge.Query) (total int, res []*dao_recharge.List, err error) {
	m := dao.SysRecharge.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysRecharge.Columns().CreateTime)
	if req.Code != "" {
		m = m.WhereIn(dao.SysRecharge.Columns().Code, req.Code)
	}

	if req.Name != "" {
		id, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Name, req.Name).Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.WhereIn(dao.SysRecharge.Columns().UserId, id)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysRecharge.Columns().Status, req.Status)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysRecharge
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_recharge.List, len(list))
	for i, v := range list {
		var entity *dao_recharge.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, v.UserId).
			Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.User = user.String()

		res[i] = entity
	}

	return
}
