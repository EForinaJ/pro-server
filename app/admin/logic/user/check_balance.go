package user

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/shopspring/decimal"
)

// CheckBalance implements service.IUser.
func (s *sUser) CheckBalance(ctx context.Context, id int64, money float64) (res bool, err error) {
	balance, err := dao.SysUser.Ctx(ctx).
		WherePri(id).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}

	return decimal.NewFromFloat(money).LessThanOrEqual(decimal.NewFromFloat(balance.Float64())), nil
}
