package distribute

import (
	"context"
	dto_distribute "server/app/admin/dto/distribute"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckCancel implements service.IDistribute.
func (s *sDistribute) CheckCancel(ctx context.Context, req *dto_distribute.Cancel) (err error) {
	obj, err := dao.SysDistribute.Ctx(ctx).Where(dao.SysDistribute.Columns().Id, req.Id).
		Fields(dao.SysDistribute.Columns().OrderId,
			dao.SysDistribute.Columns().IsCancel).One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if gconv.Int(obj.GMap().Get(dao.SysDistribute.Columns().IsCancel)) == consts.Yes {
		return utils_error.ErrMessage(response.FAILD, "派单已取消，无法重新取消派单")
	}

	orderStatus, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, obj.GMap().Get(dao.SysDistribute.Columns().OrderId)).
		Value(dao.SysOrder.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if orderStatus.Int() == consts.OrderStatusCancel {
		return utils_error.ErrMessage(response.FAILD, "订单已取消，无法取消派单")
	}
	if orderStatus.Int() == consts.OrderStatusComplete {
		return utils_error.ErrMessage(response.FAILD, "订单已完成，无法取消派单")
	}
	if orderStatus.Int() == consts.OrderStatusPendingPayment {
		return utils_error.ErrMessage(response.FAILD, "订单未支付，无法取消派单")
	}
	if orderStatus.Int() == consts.OrderStatusRefund {
		return utils_error.ErrMessage(response.FAILD, "订单已退款，无法取消派单")
	}

	return
}
