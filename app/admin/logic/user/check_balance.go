package user

import (
	"context"
	dto_user "server/app/admin/dto/user"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/shopspring/decimal"
)

// CheckBalance implements service.IUser.
func (s *sUser) CheckBalance(ctx context.Context, req *dto_user.ChangeBalance) (err error) {
	balance, err := dao.SysUser.Ctx(ctx).
		WherePri(req.Id).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if decimal.NewFromFloat(req.Amount).LessThanOrEqual(decimal.NewFromFloat(balance.Float64())) {
		return utils_error.ErrMessage(response.FAILD, "余额不足，减少金额超出余额")
	}
	return nil
}
