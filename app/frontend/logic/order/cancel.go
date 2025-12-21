package order

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// Cancel implements service.IOrder.
func (s *sOrder) Cancel(ctx context.Context, id int64) (err error) {

	_, err = dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		Where(dao.SysOrder.Columns().Status, consts.OrderStatusPendingPayment).
		Data(g.Map{
			dao.SysOrder.Columns().Status: consts.OrderStatusCancel,
		}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
