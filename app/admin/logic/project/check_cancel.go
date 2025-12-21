package project

import (
	"context"
	dto_project "server/app/admin/dto/project"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
)

// CheckCancel implements service.IProject.
func (s *sProject) CheckCancel(ctx context.Context, req *dto_project.Cancel) (err error) {
	status, err := dao.SysProject.Ctx(ctx).
		Where(dao.SysProject.Columns().Id, req.Id).
		Value(dao.SysProject.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if status.Int() == consts.ProjectStatusCancel {
		return utils_error.ErrMessage(response.FAILD, "项目已取消")
	}
	if status.Int() == consts.ProjectStatusSettlement {
		return utils_error.ErrMessage(response.FAILD, "项目已结算")
	}

	return nil
}
