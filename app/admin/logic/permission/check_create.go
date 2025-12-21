package permission

import (
	"context"
	dto_permission "server/app/admin/dto/permission"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckCreate implements service.IPermission.
func (s *sPermission) CheckCreate(ctx context.Context, req *dto_permission.Create) (err error) {
	exist, err := dao.SysPermission.Ctx(ctx).Where(dao.SysPermission.Columns().Permission, req.Permission).Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if exist {
		return utils_error.ErrMessage(response.FAILD, "权限标识已存在")
	}
	return
}
