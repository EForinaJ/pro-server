package withdraw

import (
	"context"
	dto_withdraw "server/app/admin/dto/withdraw"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// CheckApply implements service.IWithdraw.
func (s *sWithdraw) CheckApply(ctx context.Context, req *dto_withdraw.Apply) (err error) {
	withdraw, err := dao.SysWithdraw.Ctx(ctx).
		Fields(dao.SysWithdraw.Columns().WitkeyId,
			dao.SysWithdraw.Columns().Status,
			dao.SysWithdraw.Columns().Amount).
		WherePri(req.Id).One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if withdraw.GMap().Get(dao.SysWithdraw.Columns().Status) != consts.StatusApply {
		return utils_error.ErrMessage(response.FAILD, "提现申请已经审核")
	}

	commission, err := dao.SysWitkey.Ctx(ctx).
		WherePri(gconv.Int64(withdraw.GMap().Get(dao.SysWithdraw.Columns().WitkeyId))).
		Value(dao.SysWitkey.Columns().Commission)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if decimal.NewFromFloat(commission.Float64()).GreaterThanOrEqual(decimal.NewFromFloat(gconv.Float64(withdraw.GMap().Get(dao.SysWithdraw.Columns().Amount)))) {
		return utils_error.ErrMessage(response.FAILD, "余额不足")
	}

	return
}
