package user

import (
	"context"
	dto_user "server/app/admin/dto/user"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// ChangeBalance implements service.IUser.
func (s *sUser) ChangeBalance(ctx context.Context, req *dto_user.ChangeBalance) (err error) {

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
	balance, err := tx.Model(dao.SysUser.Table()).
		WherePri(req.Id).Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	amount := decimal.NewFromFloat(balance.Float64())
	entity := g.Map{
		dao.SysBalance.Columns().After:      balance,
		dao.SysBalance.Columns().Amount:     req.Amount,
		dao.SysBalance.Columns().Mode:       req.Mode,
		dao.SysBalance.Columns().UserId:     req.Id,
		dao.SysBalance.Columns().CreateTime: gtime.Now(),
		dao.SysBalance.Columns().Type:       consts.UserChangeBalanceTypeSystem,
	}

	switch req.Mode {
	case consts.FundLogModeAdd:
		newBalance := amount.Add(decimal.NewFromFloat(req.Amount))
		_, err = tx.Model(dao.SysUser.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysUser.Columns().Balance: newBalance,
		}).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		entity[dao.SysBalance.Columns().Before] = newBalance
		entity[dao.SysBalance.Columns().Remark] = "系统增加余额"
		entity[dao.SysBalance.Columns().Related] = "系统增加余额"
	case consts.FundLogModeSub:
		newBalance := amount.Sub(decimal.NewFromFloat(req.Amount))
		_, err = tx.Model(dao.SysUser.Table()).WherePri(req.Id).Data(g.Map{
			dao.SysUser.Columns().Balance: newBalance,
		}).Update()
		if err != nil {
			return utils_error.Err(response.DB_SAVE_ERROR)
		}
		entity[dao.SysBalance.Columns().Related] = "系统减少余额"
		entity[dao.SysBalance.Columns().Remark] = "系统减少余额"
		entity[dao.SysBalance.Columns().Before] = newBalance
	}

	if req.Remark != "" {
		entity[dao.SysBalance.Columns().Remark] = req.Remark
	}

	_, err = tx.Model(dao.SysBalance.Table()).
		Data(entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
