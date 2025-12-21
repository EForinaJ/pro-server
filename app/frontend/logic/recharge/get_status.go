package recharge

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetStatus implements service.IRecharge.
func (s *sRecharge) GetStatus(ctx context.Context, id int64) (status int, err error) {
	rechargeStatus, err := dao.SysRecharge.Ctx(ctx).
		Where(dao.SysRecharge.Columns().UserId, ctx.Value("userId")).
		Where(dao.SysRecharge.Columns().Id, id).
		Value(dao.SysRecharge.Columns().Status)
	if err != nil {
		return 0, utils_error.Err(response.DB_READ_ERROR)
	}
	return rechargeStatus.Int(), nil
}
