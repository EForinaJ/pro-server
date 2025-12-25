package order

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckStartService implements service.IOrder.
func (s *sOrder) CheckStartService(ctx context.Context, id int64) (err error) {

	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Fields(dao.SysOrder.Columns().WitkeyCount, dao.SysOrder.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusPendingPayment {
		return utils_error.ErrMessage(response.FAILD, "订单待支付，无法开始")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusInProgress {
		return utils_error.ErrMessage(response.FAILD, "订单进行中，无法开始")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusComplete {
		return utils_error.ErrMessage(response.FAILD, "订单已完成，无法开始")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusCancel {
		return utils_error.ErrMessage(response.FAILD, "订单已取消，无法开始")
	}
	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) == consts.OrderStatusRefund {
		return utils_error.ErrMessage(response.FAILD, "订单已退款，无法开始")
	}

	count, err := dao.SysDistribute.Ctx(ctx).
		Where(dao.SysDistribute.Columns().OrderId, id).
		Where(dao.SysDistribute.Columns().IsCancel, consts.Not).Count()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if count < gconv.Int(order.GMap().Get(dao.SysOrder.Columns().WitkeyCount)) {
		return utils_error.ErrMessage(response.FAILD, "订单服务人数未达到")
	}

	return
}
