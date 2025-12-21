package project

import (
	"context"
	dto_project "server/app/admin/dto/project"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// CheckSettlement implements service.IProject.
func (s *sProject) CheckSettlement(ctx context.Context, req *dto_project.Settlement) (err error) {

	detail, err := dao.SysProject.Ctx(ctx).
		Where(dao.SysProject.Columns().Id, req.Id).
		Fields(dao.SysProject.Columns().Status, dao.SysProject.Columns().Money).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if gconv.Int(detail.GMap().Get(dao.SysProject.Columns().Status)) == consts.ProjectStatusPending {
		return utils_error.ErrMessage(response.FAILD, "项目未开始服务")
	}
	if gconv.Int(detail.GMap().Get(dao.SysProject.Columns().Status)) == consts.ProjectStatusInProgress {
		return utils_error.ErrMessage(response.FAILD, "项目进行中")
	}
	if gconv.Int(detail.GMap().Get(dao.SysProject.Columns().Status)) == consts.ProjectStatusCancel {
		return utils_error.ErrMessage(response.FAILD, "项目已取消")
	}
	if gconv.Int(detail.GMap().Get(dao.SysProject.Columns().Status)) == consts.ProjectStatusSettlement {
		return utils_error.ErrMessage(response.FAILD, "项目已结算")
	}

	if decimal.NewFromFloat(gconv.Float64(detail.GMap().Get(dao.SysProject.Columns().Money))).LessThan(decimal.NewFromFloat(req.Money)) {
		return utils_error.ErrMessage(response.FAILD, "结算的佣金超出项目预算")
	}

	return
}
