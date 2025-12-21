package order

import (
	"context"
)

// CheckCanCreate implements service.IOrder.
func (s *sOrder) CheckCanCreate(ctx context.Context) (res bool, err error) {
	// manageId := ctx.Value("userId")
	// exist, err := dao.SysOrder.Ctx(ctx).
	// 	Where(dao.SysOrder.Columns().ManageId, manageId).
	// 	Where(dao.SysOrder.Columns().Status, consts.OrderStatusPendingPayment).
	// 	Exist()
	// if err != nil {
	// 	return false, utils_error.Err(response.DB_READ_ERROR)
	// }

	// return exist, nil
	return
}
