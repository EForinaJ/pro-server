package user

import (
	"context"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckUniquePhone implements service.IUser.
func (s *sUser) CheckUniquePhone(ctx context.Context, phone string) (res bool, err error) {
	res, err = dao.SysUser.Ctx(ctx).
		Where(dao.SysUser.Columns().Phone, phone).Exist()
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
