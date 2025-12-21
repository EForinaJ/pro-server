package role

import (
	"context"
	dto_role "server/app/admin/dto/role"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckEdit implements service.IRole.
func (s *sRole) CheckEdit(ctx context.Context, req *dto_role.Edit) (err error) {
	exist, err := dao.SysRole.Ctx(ctx).
		Where(dao.SysRole.Columns().Code, req.Code).
		WhereNotIn(dao.SysRole.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if exist {
		return utils_error.ErrMessage(response.FAILD, "角色已存在")
	}

	return
}
