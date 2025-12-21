package order

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_order "server/app/frontend/dao/order"
	dto_order "server/app/frontend/dto/order"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IOrder.
func (s *sOrder) GetList(ctx context.Context, req *dto_order.Query) (total int, res []*dao_order.List, err error) {
	m := dao.SysOrder.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		OrderDesc(dao.SysOrder.Columns().CreateTime)

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

		// 商品内容
		product, err := dao.SysProduct.Ctx(ctx).
			Where(dao.SysProduct.Columns().Id, v.ProductId).
			Fields(
				dao.SysProduct.Columns().Name,
				dao.SysProduct.Columns().Pic,
				dao.SysProduct.Columns().GameId,
				dao.SysProduct.Columns().Type,
			).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Product = &dao_order.Product{
			Id:   v.ProductId,
			Name: gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
			Pic:  gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
		}

		game, err := dao.SysGame.Ctx(ctx).
			Where(dao.SysGame.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().GameId)).
			Value(dao.SysGame.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Product.Game = game.String()

		dictData, err := dao.SysDictData.Ctx(ctx).
			Where(dao.SysDictData.Columns().Code, "PRODUCT_TYPE").
			Where(dao.SysDictData.Columns().Value, product.GMap().Get(dao.SysProduct.Columns().Type)).
			Value(dao.SysDictData.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Product.Type = dictData.String()

		res[i] = entity
	}

	return
}
