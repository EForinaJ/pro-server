package withdraw

import (
	"context"
	dao_withdraw "server/app/admin/dao/withdraw"
	dto_withdraw "server/app/admin/dto/withdraw"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IWithdraw.
func (s *sWithdraw) GetList(ctx context.Context, req *dto_withdraw.Query) (total int, res []*dao_withdraw.List, err error) {
	m := dao.SysWithdraw.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysWithdraw.Columns().CreateTime)
	if req.Code != "" {
		m = m.WhereIn(dao.SysWithdraw.Columns().Code, req.Code)
	}
	if req.WitkeyId != 0 {
		m = m.WhereIn(dao.SysWithdraw.Columns().UserId, req.WitkeyId)
	}

	if req.Phone != "" {
		// id, err := dao.SysUser.Ctx(ctx).
		// 	Where(dao.SysUser.Columns().Phone, req.Phone).Value(dao.SysUser.Columns().Id)
		// if err != nil {
		// 	return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		// }
		// m = m.Where(dao.SysWithdraw.Columns().UserId, id)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysWithdraw.Columns().Status, req.Status)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysWithdraw
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_withdraw.List, len(list))
	for i, v := range list {
		var entity *dao_withdraw.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		// user, err := dao.SysUser.Ctx(ctx).
		// 	Where(dao.SysUser.Columns().Id, v.UserId).
		// 	Value(dao.SysUser.Columns().Name)
		// if err != nil {
		// 	return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		// }
		// entity.User = user.String()

		//
		manage, err := dao.SysManage.Ctx(ctx).
			Where(dao.SysManage.Columns().Id, v.ManageId).
			Value(dao.SysManage.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Manage = gconv.String(manage)
		res[i] = entity
	}

	return
}
