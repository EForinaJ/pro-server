package account

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	utils_snow "server/app/common/utils/snow"
	dto_account "server/app/frontend/dto/account"
	"server/internal/dao"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// Withdraw implements service.IAccount.
func (s *sAccount) Withdraw(ctx context.Context, req *dto_account.Withdraw) (err error) {

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

	withdrawSetting, err := tx.Model(dao.SysConfig.Table()).
		Where(dao.SysConfig.Columns().Key, consts.WithdrawSetting).
		Value(dao.SysConfig.Columns().Value)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	withdrawSettingJson, err := gjson.DecodeToJson(withdrawSetting)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	withDrawPercentJson := withdrawSettingJson.Get("withDrawPercent")

	withDrawPercent := decimal.
		NewFromFloat(gconv.Float64(withDrawPercentJson)).Div(decimal.NewFromFloat(100))
	serviceFee := decimal.NewFromFloat(req.Money).Mul(withDrawPercent)
	settledAmount := decimal.NewFromFloat(req.Money).Sub(serviceFee)

	card, err := tx.Model(dao.SysUserCard.Table()).
		Where(dao.SysUserCard.Columns().Id, req.Id).One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	// 添加提现申请
	_, err = tx.Model(dao.SysWithdraw.Table()).Data(g.Map{
		dao.SysWithdraw.Columns().Code:          utils_snow.GetCode(ctx, consts.TX),
		dao.SysWithdraw.Columns().Name:          card.GMap().Get(dao.SysUserCard.Columns().Name),
		dao.SysWithdraw.Columns().Number:        card.GMap().Get(dao.SysUserCard.Columns().Number),
		dao.SysWithdraw.Columns().Type:          card.GMap().Get(dao.SysUserCard.Columns().Type),
		dao.SysWithdraw.Columns().SettledAmount: settledAmount,
		dao.SysWithdraw.Columns().ServiceFee:    serviceFee,
		dao.SysWithdraw.Columns().Amount:        req.Money,
		dao.SysWithdraw.Columns().WitkeyId:      ctx.Value("userId"),
		dao.SysWithdraw.Columns().Status:        consts.WithdrawStatusApply,
		dao.SysWithdraw.Columns().CreateTime:    gtime.Now(),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	balance, err := tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	newBalance := decimal.NewFromFloat(balance.Float64()).Sub(decimal.NewFromFloat(req.Money))
	_, err = tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).Data(g.Map{
		dao.SysUser.Columns().Balance: newBalance,
	}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	// 添加钱包明细
	// _, err = tx.Model(dao.SysUserFundLog.Table()).
	// 	Data(g.Map{
	// 		dao.SysUserFundLog.Columns().After:  balance,
	// 		dao.SysUserFundLog.Columns().Before: newBalance,
	// 		dao.SysUserFundLog.Columns().Money:  req.Money,
	// 		// dao.SysUserFundLog.Columns().Type:       consts.UserFundLogTypeWithdraw,
	// 		dao.SysUserFundLog.Columns().Mode:       consts.FundLogModeSub,
	// 		dao.SysUserFundLog.Columns().UserId:     ctx.Value("userId"),
	// 		dao.SysUserFundLog.Columns().CreateTime: gtime.Now(),
	// 	}).Insert()
	// if err != nil {
	// 	return utils_error.Err(response.DB_SAVE_ERROR)
	// }
	return
}
