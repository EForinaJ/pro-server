package witkey

import (
	"context"
	dto_witkey "server/app/admin/dto/witkey"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/shopspring/decimal"
)

// CheckCommission implements service.IWitkey.
func (s *sWitkey) CheckCommission(ctx context.Context, req *dto_witkey.Commission) (err error) {
	commission, err := dao.SysWitkey.Ctx(ctx).
		WherePri(req.Id).
		Value(dao.SysWitkey.Columns().Commission)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if !decimal.NewFromFloat(req.Amount).LessThanOrEqual(decimal.NewFromFloat(commission.Float64())) {
		return utils_error.ErrMessage(response.FAILD, "佣金不足，减少金额超出佣金")
	}
	return
}
