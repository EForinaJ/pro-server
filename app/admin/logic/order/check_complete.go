package order

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckComplete implements service.IOrder.
func (s *sOrder) CheckComplete(ctx context.Context, id int64) (err error) {
	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Fields(dao.SysOrder.Columns().WitkeyCount, dao.SysOrder.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusPendingPayment {
		return utils_error.ErrMessage(response.FAILD, "订单待支付，无法确认完成")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusPendingOrder {
		return utils_error.ErrMessage(response.FAILD, "订单待服务，无法确认完成")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusComplete {
		return utils_error.ErrMessage(response.FAILD, "订单已完成，无法确认完成")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusCancel {
		return utils_error.ErrMessage(response.FAILD, "订单已取消，无法确认完成")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusRefund {
		return utils_error.ErrMessage(response.FAILD, "订单已退款，无法确认完成")
	}
	// //
	// count, err := dao.SysSettlement.Ctx(ctx).
	// 	Where(dao.SysSettlement.Columns().OrderId, id).
	// 	Where(dao.SysSettlement.Columns().Status, consts.Yes).Count()
	// if err != nil {
	// 	return utils_error.Err(response.DB_READ_ERROR)
	// }

	// if count < gconv.Int(order.GMap().Get(dao.SysOrder.Columns().WitkeyCount)) {
	// 	return utils_error.ErrMessage(response.FAILD, "报单结算威客人数不足，无法确认完成")
	// }

	return
}
