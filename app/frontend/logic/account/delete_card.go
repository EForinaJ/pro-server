package account

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// DeleteCard implements service.IAccount.
func (s *sAccount) DeleteCard(ctx context.Context, id int64) (err error) {
	_, err = dao.SysUserCard.Ctx(ctx).
		Where(dao.SysUserCard.Columns().Id, id).
		Where(dao.SysUserCard.Columns().UserId, ctx.Value("userId")).Delete()
	if err != nil {
		return utils_error.Err(response.DELETE_FAILED)
	}
	return
}
