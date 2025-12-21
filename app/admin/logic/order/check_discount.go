package order

import (
	"context"
	dto_order "server/app/admin/dto/order"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// CheckDiscount implements service.IOrder.
func (s *sOrder) CheckDiscount(ctx context.Context, req *dto_order.AddDiscount) (err error) {
	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, req.Id).
		Fields(dao.SysOrder.Columns().TotalAmount, dao.SysOrder.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if !decimal.NewFromFloat(req.Money).LessThanOrEqual(decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().TotalAmount)))) {
		return utils_error.ErrMessage(response.FAILD, "优惠金额超出订单总额")
	}

	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) != consts.OrderStatusPendingPayment {
		return utils_error.ErrMessage(response.FAILD, "订单状态不为待支付，无法添加优惠")
	}

	return nil
}
