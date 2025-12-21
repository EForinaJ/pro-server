package project

import (
	"context"
	dao_project "server/app/admin/dao/project"
	dto_project "server/app/admin/dto/project"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IProject.
func (s *sProject) GetList(ctx context.Context, req *dto_project.Query) (total int, res []*dao_project.List, err error) {
	m := dao.SysProject.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysProject.Columns().CreateTime)

	if req.Code != "" {
		m = m.Where(dao.SysProject.Columns().Code, req.Code)
	}

	if req.Order != "" {
		orderCode, err := dao.SysOrder.Ctx(ctx).
			WhereLike(dao.SysOrder.Columns().Code, "%"+req.Order+"%").
			Value(dao.SysOrder.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysProject.Columns().OrderId, orderCode)
	}

	if req.Status != 0 {
		m = m.Where(dao.SysProject.Columns().Status, req.Status)
	}

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysProject
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_project.List, len(list))
	for i, v := range list {
		var entity *dao_project.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		// 接单者
		witkeyUserId, err := dao.SysWitkey.Ctx(ctx).
			Where(dao.SysWitkey.Columns().Id, v.WitkeyId).
			Value(dao.SysWitkey.Columns().UserId)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, witkeyUserId.Int64()).
			Fields(dao.SysUser.Columns().Name,
				dao.SysUser.Columns().Avatar).One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Witkey = &dao_project.Witkey{
			Id:     witkeyUserId.Int64(),
			Name:   gconv.String(user.GMap().Get(dao.SysUser.Columns().Name)),
			Avatar: gconv.String(user.GMap().Get(dao.SysUser.Columns().Avatar)),
		}

		//  获取订单号
		order, err := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Id, v.OrderId).
			Fields(dao.SysOrder.Columns().Code,
				dao.SysOrder.Columns().ProductId,
				dao.SysOrder.Columns().Specifications,
				dao.SysOrder.Columns().UnitPrice,
				dao.SysOrder.Columns().Quantity,
			).One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Order = gconv.String(order.GMap().Get(dao.SysOrder.Columns().Code))

		// 商品内容
		product, err := dao.SysProduct.Ctx(ctx).
			Where(dao.SysProduct.Columns().Id, order.GMap().Get(dao.SysOrder.Columns().ProductId)).
			Fields(
				dao.SysProduct.Columns().Name,
				dao.SysProduct.Columns().Pic,
			).
			One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Product = &dao_project.Product{
			Id:             gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().ProductId)),
			Name:           gconv.String(product.GMap().Get(dao.SysProduct.Columns().Name)),
			Pic:            gconv.String(product.GMap().Get(dao.SysProduct.Columns().Pic)),
			Specifications: gconv.String(order.GMap().Get(dao.SysOrder.Columns().Specifications)),
			Quantity:       gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Quantity)),
			UnitPrice:      gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().UnitPrice)),
		}

		res[i] = entity
	}
	return
}
