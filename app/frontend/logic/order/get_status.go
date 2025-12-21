package order

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// GetStatus implements service.IOrder.
func (s *sOrder) GetStatus(ctx context.Context, id int64) (status int, err error) {
	orderStatus, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		Where(dao.SysOrder.Columns().Id, id).
		Value(dao.SysOrder.Columns().Status)
	if err != nil {
		return 0, utils_error.Err(response.DB_READ_ERROR)
	}
	return orderStatus.Int(), nil
}
