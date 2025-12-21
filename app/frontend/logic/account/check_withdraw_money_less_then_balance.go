package account

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"

	"github.com/shopspring/decimal"
)

// CheckWithdrawMoneyLessthanBalance implements service.IAccount.
func (s *sAccount) CheckWithdrawMoneyLessthanBalance(ctx context.Context, req *dto_account.Withdraw) (err error) {
	balance, err := dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if decimal.NewFromFloat(balance.Float64()).LessThan(decimal.NewFromFloat(req.Money)) {
		return utils_error.ErrMessage(response.FAILD, "用户余额不足")
	}
	return nil
}
