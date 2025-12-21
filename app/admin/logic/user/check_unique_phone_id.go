package user

import (
	"context"
	dto_user "server/app/admin/dto/user"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckUniquePhoneId implements service.IUser.
func (s *sUser) CheckUniquePhoneId(ctx context.Context, req *dto_user.Edit) (res bool, err error) {
	res, err = dao.SysUser.Ctx(ctx).
		WhereNotIn(dao.SysUser.Columns().Id, req.Id).
		Where(dao.SysUser.Columns().Phone, req.Phone).Exist()
	if err != nil {
		return false, utils_error.Err(response.DB_READ_ERROR)
	}
	return
}
