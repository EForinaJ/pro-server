package order

import (
	"context"
	dao_order "server/app/admin/dao/order"
	dto_order "server/app/admin/dto/order"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetWitkeyList implements service.IOrder.
func (s *sOrder) GetWitkeyList(ctx context.Context, req *dto_order.WitkeyQuery) (total int, res []*dao_order.WitkeyList, err error) {
	m := dao.SysOrderWitkey.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysOrderWitkey.Columns().CreateTime)
	m = m.WhereIn(dao.SysOrderWitkey.Columns().OrderId, req.Id)
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysOrderWitkey
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_order.WitkeyList, len(list))
	for i, v := range list {
		var entity *dao_order.WitkeyList
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		witkey, err := dao.SysWitkey.
			Ctx(ctx).Where(dao.SysWitkey.Columns().Id, v.WitkeyId).
			Fields(
				dao.SysWitkey.Columns().UserId,
				dao.SysWitkey.Columns().GameId,
				dao.SysWitkey.Columns().Rate,
				dao.SysWitkey.Columns().TitleId).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		userName, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, witkey.GMap().Get(dao.SysWitkey.Columns().UserId)).
			Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Name = userName.String()

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id, witkey.GMap().Get(dao.SysWitkey.Columns().GameId)).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Game = game.String()

		title, err := dao.SysTitle.Ctx(ctx).
			Where(dao.SysTitle.Columns().Id, witkey.GMap().Get(dao.SysWitkey.Columns().TitleId)).
			Value(dao.SysTitle.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Title = title.String()

		res[i] = entity
	}

	return
}
