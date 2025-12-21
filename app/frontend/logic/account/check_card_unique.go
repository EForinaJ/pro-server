package account

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"
)

// CheckCardUnique implements service.IAccount.
func (s *sAccount) CheckCardUnique(ctx context.Context, req *dto_account.CreateCard) (err error) {
	exist, err := dao.SysUserCard.Ctx(ctx).
		Where(dao.SysUserCard.Columns().UserId, ctx.Value("userId")).
		Where(dao.SysUserCard.Columns().Type, req.Type).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	if exist {
		return utils_error.ErrMessage(response.FAILD, "该类型提现账户已经创建过了")
	}
	return nil
}
