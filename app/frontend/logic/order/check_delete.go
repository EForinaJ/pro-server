package order

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckDelete implements service.IOrder.
func (s *sOrder) CheckDelete(ctx context.Context, id int64) (err error) {
	order, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		Fields(dao.SysOrder.Columns().Status).
		One()

	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if order.IsEmpty() {
		return utils_error.ErrMessage(response.FAILD, "该订单不是你的")
	}

	if gconv.Int(order.GMap().Get(dao.SysOrder.Columns().Status)) != consts.OrderStatusCancel {
		return utils_error.ErrMessage(response.FAILD, "该订单未取消，无法删除")
	}

	return
}
