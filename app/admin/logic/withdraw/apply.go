package withdraw

import (
	"context"
	dto_withdraw "server/app/admin/dto/withdraw"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
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
	_, err = tx.Model(dao.SysWithdraw.Table()).
		Where(dao.SysWithdraw.Columns().Id, req.Id).Data(g.Map{
		dao.SysWithdraw.Columns().ReceiptNum:   req.ReceiptNum,
		dao.SysWithdraw.Columns().ReceiptFiles: req.ReceiptFile,
		dao.SysWithdraw.Columns().Status:       req.Status,
		dao.SysWithdraw.Columns().Remark:       req.Remark,
		dao.SysWithdraw.Columns().ManageId:     ctx.Value("userId"),
	}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	if req.Status == consts.WithdrawStatuSuccess {

		withdraw, err := tx.Model(dao.SysWithdraw.Table()).
			Where(dao.SysWithdraw.Columns().Id, req.Id).
			One()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		money := decimal.
			NewFromFloat(gconv.Float64(withdraw.GMap().Get(dao.SysWithdraw.Columns().Money)))

		// 修改佣金
		userCommission, err := tx.Model(dao.SysUser.Table()).
			Where(dao.SysUser.Columns().Id,
				withdraw.GMap().Get(dao.SysWithdraw.Columns().UserId)).Value()
		// Value(dao.SysUser.Columns().Commission)
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		amount := decimal.NewFromFloat(userCommission.Float64()).Sub(money)

		_, err = tx.Model(dao.SysUser.Table()).
			Where(dao.SysUser.Columns().Id, withdraw.GMap().Get(dao.SysWithdraw.Columns().UserId)).
			Data(g.Map{
				// dao.SysUser.Columns().Commission: amount,
			}).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

		_, err = tx.Model(dao.SysBalance.Table()).
			Data(g.Map{
				dao.SysBalance.Columns().After:  userCommission,
				dao.SysBalance.Columns().Before: amount,
				dao.SysBalance.Columns().Money:  money,
				// dao.SysUserFundLog.Columns().Type:       consts.UserFundLogTypeWithdraw,
				dao.SysBalance.Columns().Mode:       consts.FundLogModeSub,
				dao.SysBalance.Columns().UserId:     withdraw.GMap().Get(dao.SysWithdraw.Columns().UserId),
				dao.SysBalance.Columns().CreateTime: gtime.Now(),
			}).Insert()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}

	}

	return
}
