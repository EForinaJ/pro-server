package aftersales

import (
	"context"
	dao_aftersales "server/app/admin/dao/aftersales"
	dto_aftersales "server/app/admin/dto/aftersales"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetList implements service.IAftersales.
func (s *sAftersales) GetList(ctx context.Context, req *dto_aftersales.Query) (total int, res []*dao_aftersales.List, err error) {
	m := dao.SysAftersales.Ctx(ctx).
		Page(req.Page, req.Limit).
		OrderDesc(dao.SysAftersales.Columns().CreateTime)
	if req.OrderCode != "" {
		orderId, err := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Code, req.OrderCode).Value(dao.SysOrder.Columns().Id)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		m = m.Where(dao.SysAftersales.Columns().OrderId, orderId)
	}

	if req.Code != "" {
		m = m.Where(dao.SysAftersales.Columns().Code, req.Code)
	}
	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysAftersales
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_aftersales.List, len(list))
	for i, v := range list {
		var entity *dao_aftersales.List
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		//  订单编号
		order, err := dao.SysOrder.Ctx(ctx).
			Where(dao.SysOrder.Columns().Id, v.OrderId).
			Fields(dao.SysOrder.Columns().Code, dao.SysOrder.Columns().ProductId).One()
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.OrderCode = gconv.String(order.GMap().Get(dao.SysOrder.Columns().Code))

		// 订单商品
		var product *dao_aftersales.Product
		err = dao.SysProduct.Ctx(ctx).
			Where(dao.SysProduct.Columns().Id, order.GMap().Get(dao.SysOrder.Columns().ProductId)).
			Fields(dao.SysProduct.Columns().Id, dao.SysProduct.Columns().Name, dao.SysProduct.Columns().Pic).Scan(&product)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Product = product
		//  申请用户
		user, err := dao.SysUser.Ctx(ctx).
			Where(dao.SysUser.Columns().Id, v.UserId).
			Value(dao.SysUser.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.User = user.String()

		res[i] = entity
	}

	return
}
