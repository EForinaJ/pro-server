package withdraw

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// CheckCommission implements service.IWithdraw.
func (s *sWithdraw) CheckCommission(ctx context.Context, id int64) (res bool, err error) {
	withdraw, err := dao.SysWithdraw.Ctx(ctx).
		Fields(dao.SysWithdraw.Columns().WitkeyId,
			dao.SysWithdraw.Columns().Amount).
		WherePri(id).One()
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}

	commission, err := dao.SysUser.Ctx(ctx).
		WherePri(gconv.Int64(withdraw.GMap().
			Get(dao.SysWithdraw.Columns().WitkeyId))).Value()
	// Value(dao.SysUser.Columns().Commission)
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}

	return decimal.NewFromFloat(commission.Float64()).GreaterThanOrEqual(decimal.NewFromFloat(gconv.Float64(withdraw.GMap().Get(dao.SysWithdraw.Columns().Amount)))), nil
}
