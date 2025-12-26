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
		m = m.Where(dao.SysWithdraw.Columns().Code, req.Code)
	}
	if req.Name != "" {
		userId, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Name, req.Name).
			Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		witkeyId, err := dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().UserId, userId).
			Value(dao.SysWitkey.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysWithdraw.Columns().WitkeyId, witkeyId)
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

		userId, err := dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().Id, v.WitkeyId).
			Value(dao.SysWitkey.Columns().UserId)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, userId).
			Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Witkey = user.String()
		res[i] = entity
	}

	return
}
