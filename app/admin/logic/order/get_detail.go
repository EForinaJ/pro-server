package order

import (
	"context"
	dao_order "server/app/admin/dao/order"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
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

	user, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, order.GMap().Get(dao.SysOrder.Columns().UserId)).
		Fields(dao.SysUser.Columns().Name,
			dao.SysUser.Columns().Avatar).
		One()
	if err != nil {
		return nil, utils_error.Err(response.FAILD)
	}

	orderInfo.User = &dao_order.User{
		Id:     gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().UserId)),
		Name:   gconv.String(user.GMap().Get(dao.SysUser.Columns().Name)),
		Avatar: gconv.String(user.GMap().Get(dao.SysUser.Columns().Avatar)),
	}

	product, err := dao.SysProduct.Ctx(ctx).
		Where(dao.SysProduct.Columns().Id,
			order.GMap().Get(dao.SysOrder.Columns().ProductId),
		).
		Fields(
			dao.SysProduct.Columns().Name,
			dao.SysProduct.Columns().Pic,
			dao.SysProduct.Columns().GameId,
			dao.SysProduct.Columns().CategoryId,
		).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	orderInfo.Product = &dao_order.Product{
		Id:   gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().ProductId)),
		Name: gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
		Pic:  gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
		// Specifications: gconv.String(order.GMap().Get(dao.SysOrder.Columns().Specifications)),
		Quantity:  gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Quantity)),
		UnitPrice: gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().UnitPrice)),
		Unit:      gconv.String(order.GMap().Get(dao.SysOrder.Columns().Unit)),
	}

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	orderInfo.Product.Game = game.String()

	category, err := dao.SysCategory.Ctx(ctx).
		Where(dao.SysCategory.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().CategoryId)).
		Value(dao.SysCategory.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	orderInfo.Product.Category = category.String()

	return orderInfo, nil
}
