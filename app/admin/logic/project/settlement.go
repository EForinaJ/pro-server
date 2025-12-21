package project

import (
	"context"
	dto_project "server/app/admin/dto/project"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	utils_snow "server/app/common/utils/snow"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// Settlement implements service.IProject.
func (s *sProject) Settlement(ctx context.Context, req *dto_project.Settlement) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	detail, err := tx.Model(dao.SysProject.Table()).WherePri(req.Id).One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	//  获取配置
	titleId, err := tx.Model(dao.SysWitkey.Table()).
		WherePri(detail.GMap().Get(dao.SysProject.Columns().WitkeyId)).
		Value((dao.SysWitkey.Columns().TitleId))
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	titleServicePercent, err := dao.SysTitle.Ctx(ctx).
		WherePri(titleId).
		Value(dao.SysTitle.Columns().ServicePercent)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	//  计算项目费用
	servicePercent := decimal.NewFromFloat(titleServicePercent.Float64()).
		Div(decimal.NewFromFloat(100))
	projectMoney := decimal.NewFromFloat(req.Money)

	projectCommission := projectMoney.Sub(projectMoney.Mul(servicePercent))
	projectServiceFee := projectMoney.Sub(projectCommission)
	_, err = tx.Model(dao.SysProject.Table()).
		Where(dao.SysProject.Columns().Id, req.Id).
		Update(g.Map{
			dao.SysProject.Columns().Status:     consts.ProjectStatusSettlement,
			dao.SysProject.Columns().ServiceFee: projectServiceFee,
			dao.SysProject.Columns().Commission: projectCommission,
		})
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	witkeyUser, err := tx.Model(dao.SysWitkey.Table()).
		Where(dao.SysWitkey.Columns().Id,
			detail.GMap().Get(dao.SysProject.Columns().WitkeyId)).
		Value(dao.SysWitkey.Columns().UserId)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	userBalance, err := tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id,
			witkeyUser.Int64()).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	newBalance := decimal.NewFromFloat(userBalance.Float64()).
		Add(decimal.NewFromFloat(gconv.Float64(projectCommission)))

	_, err = tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, witkeyUser.Int64()).
		Update(g.Map{
			dao.SysUser.Columns().Balance: newBalance,
		})
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	// 添加项目日志
	_, err = tx.Model(dao.SysProjectLog.Table()).Data(g.Map{
		dao.SysProjectLog.Columns().ProjectId:  req.Id,
		dao.SysProjectLog.Columns().CreateTime: gtime.Now(),
		dao.SysProjectLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysProjectLog.Columns().Type:       consts.ProjectLogTypeSettlement,
		dao.SysProjectLog.Columns().Content:    "结算佣金：" + gconv.String(projectCommission),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	billEntity := g.Map{
		dao.SysUserBill.Columns().UserId:     witkeyUser.Int64(),
		dao.SysUserBill.Columns().RelatedId:  req.Id,
		dao.SysUserBill.Columns().Code:       utils_snow.GetCode(ctx, consts.BL),
		dao.SysUserBill.Columns().Type:       consts.BillTypeSettlementCommission,
		dao.SysUserBill.Columns().Money:      req.Money,
		dao.SysUserBill.Columns().Mode:       consts.Add,
		dao.SysUserBill.Columns().CreateTime: gtime.Now(),
	}
	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(billEntity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
