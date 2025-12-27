package user

import (
	"context"
	dto_user "server/app/admin/dto/user"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	utils_snow "server/app/common/utils/snow"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// Recharge implements service.IUser.
func (s *sUser) Recharge(ctx context.Context, req *dto_user.Recharge) (err error) {
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
	code := utils_snow.GetCode(ctx, consts.CZ)
	rs, err := tx.Model(dao.SysRecharge.Table()).
		Data(g.Map{
			dao.SysRecharge.Columns().Code:       code,
			dao.SysRecharge.Columns().Amount:     req.Amount,
			dao.SysRecharge.Columns().PayMode:    consts.PayModePersonalTransfer,
			dao.SysRecharge.Columns().UserId:     req.Id,
			dao.SysRecharge.Columns().Status:     consts.PayStatusSuccess,
			dao.SysRecharge.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	rid, err := rs.LastInsertId()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	userBalance, err := tx.Model(dao.SysUser.Table()).
		WherePri(req.Id).Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	balance := decimal.NewFromFloat(userBalance.Float64())
	newBalance := balance.Add(decimal.NewFromFloat(req.Amount))
	_, err = tx.Model(dao.SysUser.Table()).WherePri(req.Id).Data(g.Map{
		dao.SysUser.Columns().Balance: newBalance,
	}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	billEntity := g.Map{
		dao.SysUserBill.Columns().UserId:     req.Id,
		dao.SysUserBill.Columns().RelatedId:  rid,
		dao.SysUserBill.Columns().Code:       utils_snow.GetCode(ctx, consts.BL),
		dao.SysUserBill.Columns().Type:       consts.BillTypeRecharge,
		dao.SysUserBill.Columns().Amount:     req.Amount,
		dao.SysUserBill.Columns().Mode:       consts.Add,
		dao.SysUserBill.Columns().CreateTime: gtime.Now(),
	}
	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(billEntity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	//  添加日志
	_, err = tx.Model(dao.SysBalance.Table()).Data(g.Map{
		dao.SysBalance.Columns().After:      balance,
		dao.SysBalance.Columns().Amount:     req.Amount,
		dao.SysBalance.Columns().Before:     newBalance,
		dao.SysBalance.Columns().Mode:       consts.Add,
		dao.SysBalance.Columns().UserId:     req.Id,
		dao.SysBalance.Columns().CreateTime: gtime.Now(),
		dao.SysBalance.Columns().Type:       consts.UserChangeBalanceTypeRecharge,
		dao.SysBalance.Columns().Remark:     "用户充值余额",
		dao.SysBalance.Columns().Related:    code,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	//  添加支付日志
	_, err = tx.Model(dao.SysCapital.Table()).Data(g.Map{
		dao.SysCapital.Columns().CreateTime: gtime.Now(),
		dao.SysCapital.Columns().Code:       utils_snow.GetCode(ctx, consts.PM),
		dao.SysCapital.Columns().Related:    code,
		dao.SysCapital.Columns().Amount:     req.Amount,
		dao.SysCapital.Columns().Type:       consts.CapitalPaymentRecharge,
		dao.SysCapital.Columns().Mode:       consts.PayModePersonalTransfer,
		dao.SysCapital.Columns().UserId:     req.Id,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
