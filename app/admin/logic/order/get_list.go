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

// GetList implements service.IOrder.
func (s *sOrder) GetList(ctx context.Context, req *dto_order.Query) (total int, res []*dao_order.List, err error) {
	m := dao.SysOrder.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysOrder.Columns().CreateTime)
	if req.Name != "" {
		userId, err := dao.SysUser.Ctx(ctx).
			WhereLike(dao.SysUser.Columns().Name, "%"+req.Name+"%").Value(dao.SysUser.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysOrder.Columns().UserId, userId)
	}
	if req.Code != "" {
		m = m.Where(dao.SysOrder.Columns().Code, req.Code)
	}
	if req.Status != 0 {
		m = m.Where(dao.SysOrder.Columns().Status, req.Status)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysOrder
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_order.List, len(list))
	for i, v := range list {
		var entity *dao_order.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		//  下单用户
		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, v.UserId).
			Fields(dao.SysUser.Columns().Name,
				dao.SysUser.Columns().Avatar).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		entity.User = &dao_order.User{
			Id:     v.UserId,
			Name:   gconv.String(user.GMap().Get(dao.SysUser.Columns().Name)),
			Avatar: gconv.String(user.GMap().Get(dao.SysUser.Columns().Avatar)),
		}
		// 商品内容
		product, err := dao.SysProduct.Ctx(ctx).
			Where(dao.SysProduct.Columns().Id, v.ProductId).
			Fields(
				dao.SysProduct.Columns().Name,
				dao.SysProduct.Columns().Pic,
				dao.SysProduct.Columns().GameId,
			).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id,
				product.GMap().Get(dao.SysProduct.Columns().GameId)).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		entity.Product = &dao_order.Product{
			Id:        v.ProductId,
			Name:      gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
			Pic:       gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
			Quantity:  v.Quantity,
			UnitPrice: v.UnitPrice,
			Game:      game.String(),
		}

		res[i] = entity
	}
	return
}
