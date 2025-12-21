package order

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/shopspring/decimal"
)

// CheckBalanceSufficient implements service.IOrder.
func (s *sOrder) CheckBalanceSufficient(ctx context.Context, id int64) (err error) {
	actualAmount, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, id).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		Value(dao.SysOrder.Columns().ActualAmount)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	amount := decimal.NewFromFloat(actualAmount.Float64())
	userBalance, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	if decimal.NewFromFloat(userBalance.Float64()).LessThan(amount) {
		return utils_error.ErrMessage(response.PAY_ERROR, "余额不足，请充值")
	}
	return
}
