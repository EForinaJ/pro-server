package manage

import (
	"context"
	dto_manage "server/app/admin/dto/manage"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckPhoneUnique implements service.IManage.
func (s *sManage) CheckPhoneUnique(ctx context.Context, req *dto_manage.Create) (err error) {
	eixst, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Phone, req.Phone).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if eixst {
		return utils_error.ErrMessage(response.FAILD, "手机号已存在")
	}
	return
}
