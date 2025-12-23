package settlement

import (
	"context"
	dao_settlement "server/app/admin/dao/settlement"
	dto_settlement "server/app/admin/dto/settlement"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.ISettlement.
func (s *sSettlement) GetList(ctx context.Context, req *dto_settlement.Query) (total int, res []*dao_settlement.List, err error) {
	m := dao.SysSettlement.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysSettlement.Columns().CreateTime)
	if req.Name != "" {
		userId, err := dao.SysUser.Ctx(ctx).
			WhereLike(dao.SysUser.Columns().Name, "%"+req.Name+"%").Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		witkeyId, err := dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().UserId, userId).Value(dao.SysWitkey.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysSettlement.Columns().WitkeyId, witkeyId)
	}
	if req.Code != "" {
		orderId, err := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Code, req.Code).Value(dao.SysOrder.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysSettlement.Columns().OrderId, orderId)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysSettlement
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_settlement.List, len(list))
	for i, v := range list {
		var entity *dao_settlement.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		//  订单编号
		orderCode, err := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Id, v.OrderId).
			Value(dao.SysOrder.Columns().Code)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Code = orderCode.String()

		//  威客
		witkey, err := dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().Id, v.WitkeyId).
			Fields(dao.SysWitkey.Columns().UserId, dao.SysWitkey.Columns().TitleId, dao.SysWitkey.Columns().GameId).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id,
				witkey.GMap().Get(dao.SysWitkey.Columns().UserId)).
			Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Witkey = user.String()

		manage, err := dao.SysManage.Ctx(ctx).
			Where(dao.SysManage.Columns().Id,
				v.ManageId).
			Value(dao.SysManage.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Manage = manage.String()

		res[i] = entity
	}
	return
}
