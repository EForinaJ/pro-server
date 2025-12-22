package order

import (
	"context"
	dto_order "server/app/admin/dto/order"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckDistribute implements service.IOrder.
func (s *sOrder) CheckDistribute(ctx context.Context, req *dto_order.Distribute) (err error) {
	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Fields(dao.SysOrder.Columns().Status, dao.SysOrder.Columns().WitkeyCount).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) != consts.OrderStatusPendingOrder {
		return utils_error.ErrMessage(response.FAILD, "订单状态不为待服务，无法派单")
	}

	witkeyExist, err := dao.SysWitkey.Ctx(ctx).
		Where(dao.SysWitkey.Columns().Id, req.WitkeyId).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if !witkeyExist {
		return utils_error.ErrMessage(response.FAILD, "威客不存在，未找到")
	}

	count, err := dao.SysOrderWitkey.Ctx(ctx).
		Where(dao.SysOrderWitkey.Columns().OrderId, req.Id).
		Count()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if count >= gconv.Int(order.GMap().Get(dao.SysOrder.Columns().WitkeyCount)) {
		return utils_error.ErrMessage(response.FAILD, "该订单派发威客已满")
	}

	exist, err := dao.SysOrderWitkey.Ctx(ctx).
		Where(dao.SysOrderWitkey.Columns().OrderId, req.Id).
		Where(dao.SysOrderWitkey.Columns().WitkeyId, req.WitkeyId).
		Where(dao.SysOrderWitkey.Columns().IsReplaced, consts.Not).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if exist {
		return utils_error.ErrMessage(response.FAILD, "该订单以派发过该威客了")
	}

	return nil
}
