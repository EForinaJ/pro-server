package order

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckCancel implements service.IOrder.
func (s *sOrder) CheckCancel(ctx context.Context, id int64) (err error) {
	status, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Value(dao.SysOrder.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if status.Int() != consts.OrderStatusPendingPayment {
		return utils_error.ErrMessage(response.FAILD, "订单状态不为待支付，无法取消订单")
	}
	return
}
