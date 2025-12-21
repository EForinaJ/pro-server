package project

import (
	"context"
	dao_project "server/app/admin/dao/project"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetDetail implements service.IOrder.
func (s *sProject) GetDetail(ctx context.Context, id int64) (res *dao_project.Detail, err error) {

	detail, err := dao.SysProject.Ctx(ctx).
		WherePri(id).
		One()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	if detail.IsEmpty() {
		return nil, utils_error.Err(response.NOT_FOUND)
	}

	// 2. 使用结构体转换代替手动映射
	var info *dao_project.Detail
	if err := gconv.Scan(detail.Map(), &info); err != nil {
		return nil, utils_error.Err(response.UNKNOWN)
	}

	images := detail.GMap().Get(dao.SysProject.Columns().Images)
	if images != nil {
		info.Images = gconv.Strings(images)
	} else {
		info.Images = []string{}
	}

	witkeyUser, err := dao.SysWitkey.Ctx(ctx).
		WherePri(detail.GMap().Get(dao.SysProject.Columns().WitkeyId)).
		Value(dao.SysWitkey.Columns().UserId)
	if err != nil {
		return nil, utils_error.Err(response.UNKNOWN)
	}
	user, err := dao.SysUser.Ctx(ctx).
		WherePri(witkeyUser.Int64()).
		Fields(dao.SysUser.Columns().Name, dao.SysUser.Columns().Avatar).
		One()
	if err != nil {
		return nil, utils_error.Err(response.UNKNOWN)
	}
	info.Witkey = &dao_project.Witkey{
		Id:     gconv.Int64(detail.GMap().Get(dao.SysProject.Columns().WitkeyId)),
		Name:   gconv.String(user.GMap().Get(dao.SysUser.Columns().Name)),
		Avatar: gconv.String(user.GMap().Get(dao.SysUser.Columns().Avatar)),
	}

	order, err := dao.SysOrder.Ctx(ctx).
		WherePri(detail.GMap().Get(dao.SysProject.Columns().OrderId)).
		Fields(dao.SysOrder.Columns().Code,
			dao.SysOrder.Columns().ProductId,
			dao.SysOrder.Columns().Quantity,
			dao.SysOrder.Columns().UnitPrice,
			dao.SysOrder.Columns().Specifications,
			dao.SysOrder.Columns().Requirements,
			dao.SysOrder.Columns().Status).
		One()
	if err != nil {
		return nil, utils_error.Err(response.UNKNOWN)
	}
	info.Order = &dao_project.Order{
		Id:           gconv.Int64(detail.GMap().Get(dao.SysProject.Columns().OrderId)),
		Code:         gconv.String(order.GMap().Get(dao.SysOrder.Columns().Code)),
		Status:       gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)),
		Requirements: gconv.String(order.GMap().Get(dao.SysOrder.Columns().Requirements)),
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

	info.Order.Product = &dao_project.Product{
		Id:             gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().ProductId)),
		Name:           gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
		Pic:            gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
		Specifications: gconv.String(order.GMap().Get(dao.SysOrder.Columns().Specifications)),
		Quantity:       gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Quantity)),
		UnitPrice:      gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().UnitPrice)),
	}

	game, err := dao.SysGame.Ctx(ctx).
		Where(dao.SysGame.Columns().Id, product.GMap().Get(dao.SysProduct.Columns().GameId)).
		Value(dao.SysGame.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	info.Order.Product.Game = game.String()

	dictData, err := dao.SysDictData.Ctx(ctx).
		Where(dao.SysDictData.Columns().Code, "PRODUCT_TYPE").
		Where(dao.SysDictData.Columns().Value, product.GMap().Get(dao.SysProduct.Columns().Type)).
		Value(dao.SysDictData.Columns().Name)
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	info.Order.Product.Type = dictData.String()

	return info, nil
}
