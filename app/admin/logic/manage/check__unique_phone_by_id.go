package manage

import (
	"context"
	dto_manage "server/app/admin/dto/manage"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckPhoneUniqueById implements service.IManage.
func (s *sManage) CheckPhoneUniqueById(ctx context.Context, req *dto_manage.Edit) (err error) {
	exist, err := dao.SysManage.Ctx(ctx).
		Where(dao.SysManage.Columns().Phone, req.Phone).
		WhereNotIn(dao.SysManage.Columns().Id, req.Id).
		Exist()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if exist {
		return utils_error.ErrMessage(response.FAILD, "手机号已存在")
	}

	return
}
