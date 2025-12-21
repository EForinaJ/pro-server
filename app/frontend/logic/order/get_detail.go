package order

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dao_order "server/app/frontend/dao/order"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IOrder.
func (s *sOrder) GetDetail(ctx context.Context, id int64) (res *dao_order.Detail, err error) {

	order, err := dao.SysOrder.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if order.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}

	// 2. 使用结构体转换代替手动映射
	var orderInfo *dao_order.Detail
	if err := gconv.Scan(order.Map(), &orderInfo); err != nil {
		return nil, utils_error.Err(response.FAILD)
	}

	product, err := dao.SysProduct.Ctx(ctx).
		Where(dao.SysProduct.Columns().Id,
			order.GMap().Get(dao.SysOrder.Columns().ProductId),
		).
		Fields(
			dao.SysProduct.Columns().Name,
			dao.SysProduct.Columns().Pic,
			dao.SysProduct.Columns().GameId,
			dao.SysProduct.Columns().Type,
		).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	orderInfo.Product = &dao_order.Product{
		Id:   gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().ProductId)),
		Name: gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
		Pic:  gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
	}

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	orderInfo.Product.Game = game.String()

	dictData, err := dao.SysDictData.Ctx(ctx).
		Where(dao.SysDictData.Columns().Code, "PRODUCT_TYPE").
		Where(dao.SysDictData.Columns().Value, product.GMap().Get(dao.SysProduct.Columns().Type)).
		Value(dao.SysDictData.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	orderInfo.Product.Type = dictData.String()

	return orderInfo, nil
}
