package withdraw

import (
	"context"
	dto_withdraw "server/app/admin/dto/withdraw"
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

// Apply implements service.IWithdraw.
func (s *sWithdraw) Apply(ctx context.Context, req *dto_withdraw.Apply) (err error) {
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

	if req.Status == consts.StatusSuccess {
		_, err = tx.Model(dao.SysWithdraw.Table()).
			Where(dao.SysWithdraw.Columns().Id, req.Id).Data(g.Map{
			dao.SysWithdraw.Columns().Status:   consts.StatusSuccess,
			dao.SysWithdraw.Columns().ManageId: ctx.Value("userId"),
		}).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		withdraw, err := tx.Model(dao.SysWithdraw.Table()).
			Where(dao.SysWithdraw.Columns().Id, req.Id).
			One()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		amount := decimal.
			NewFromFloat(gconv.Float64(withdraw.GMap().Get(dao.SysWithdraw.Columns().Amount)))

		// 修改佣金
		commission, err := tx.Model(dao.SysWitkey.Table()).
			Where(dao.SysWitkey.Columns().Id,
				withdraw.GMap().Get(dao.SysWithdraw.Columns().WitkeyId)).
			Value(dao.SysWitkey.Columns().Commission)
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		newCommission := decimal.NewFromFloat(commission.Float64()).Sub(amount)
		_, err = tx.Model(dao.SysUser.Table()).
			Where(dao.SysUser.Columns().Id, withdraw.GMap().Get(dao.SysWithdraw.Columns().WitkeyId)).
			Data(g.Map{
				dao.SysUser.Columns().Commission: newCommission,
			}).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		_, err = tx.Model(dao.SysCommission.Table()).Data(g.Map{
			dao.SysCommission.Columns().WitkeyId:   withdraw.GMap().Get(dao.SysWithdraw.Columns().WitkeyId),
			dao.SysCommission.Columns().Before:     commission,
			dao.SysCommission.Columns().After:      newCommission,
			dao.SysCommission.Columns().Amount:     amount,
			dao.SysCommission.Columns().Mode:       consts.Sub,
			dao.SysCommission.Columns().CreateTime: gtime.Now(),
			dao.SysCommission.Columns().Type:       consts.WitkeyChangeCommissionTypeWithdraw,
			dao.SysCommission.Columns().Remark:     "威客提现变更佣金",
			dao.SysCommission.Columns().Related:    withdraw.GMap().Get(dao.SysWithdraw.Columns().Code),
		}).Insert()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		userId, err := tx.Model(dao.SysWitkey.Table()).
			Where(dao.SysWitkey.Columns().Id, withdraw.GMap().Get(dao.SysWithdraw.Columns().WitkeyId)).
			Value(dao.SysWitkey.Columns().UserId)
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		//  添加支付日志
		_, err = tx.Model(dao.SysCapital.Table()).Data(g.Map{
			dao.SysCapital.Columns().CreateTime: gtime.Now(),
			dao.SysCapital.Columns().Code:       utils_snow.GetCode(ctx, consts.PM),
			dao.SysCapital.Columns().Related:    withdraw.GMap().Get(dao.SysWithdraw.Columns().Code),
			dao.SysCapital.Columns().Amount:     amount,
			dao.SysCapital.Columns().Type:       consts.CapitalWithdrawCommission,
			dao.SysCapital.Columns().Mode:       consts.PayModePersonalTransfer,
			dao.SysCapital.Columns().UserId:     userId,
		}).Insert()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		_, err = tx.Model(dao.SysUserBill.Table()).
			Data(g.Map{
				dao.SysUserBill.Columns().UserId:     userId,
				dao.SysUserBill.Columns().RelatedId:  req.Id,
				dao.SysUserBill.Columns().Code:       utils_snow.GetCode(ctx, consts.BL),
				dao.SysUserBill.Columns().Type:       consts.BillTypeWithdrawCommission,
				dao.SysUserBill.Columns().Amount:     amount,
				dao.SysUserBill.Columns().Mode:       consts.Sub,
				dao.SysUserBill.Columns().CreateTime: gtime.Now(),
			}).Insert()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
	}
	// 拒绝结算
	if req.Status == consts.StatusFail {
		_, err = tx.Model(dao.SysWithdraw.Table()).
			Where(dao.SysWithdraw.Columns().Id, req.Id).Data(g.Map{
			dao.SysWithdraw.Columns().Status:   consts.StatusFail,
			dao.SysWithdraw.Columns().Reason:   req.Reason,
			dao.SysWithdraw.Columns().ManageId: ctx.Value("userId"),
		}).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
	}

	return
}
