package recharge

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// Revoke implements service.IRecharge.
func (s *sRecharge) Revoke(ctx context.Context, ids []int64) (err error) {
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

	userRechargeList, err := tx.Model(dao.SysRecharge.Table()).
		WhereIn(dao.SysRecharge.Columns().Id, ids).Fields(
		dao.SysRecharge.Columns().UserId,
		dao.SysRecharge.Columns().Money,
	).All()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	for _, v := range userRechargeList {
		if gconv.Float64(v.GMap().Get(dao.SysRecharge.Columns().Money)) != 0 {
			userBalance, err := tx.Model(dao.SysUser.Table()).
				WherePri(v.GMap().Get(dao.SysRecharge.Columns().UserId)).
				Value(dao.SysUser.Columns().Balance)
			if err != nil {
				return utils_error.Err(response.DB_READ_ERROR)
			}
			balance := decimal.NewFromFloat(userBalance.Float64())
			entity := g.Map{
				dao.SysBalance.Columns().After:  balance,
				dao.SysBalance.Columns().Money:  v.GMap().Get(dao.SysRecharge.Columns().Money),
				dao.SysBalance.Columns().Mode:   consts.Sub,
				dao.SysBalance.Columns().UserId: v.GMap().Get(dao.SysRecharge.Columns().UserId),
				// dao.SysUserFundLog.Columns().Type:       consts.UserFundLogTypeRevokeRecharge,
				dao.SysBalance.Columns().CreateTime: gtime.Now(),
			}
			newbalance := balance.Sub(decimal.NewFromFloat(gconv.Float64(v.GMap().Get(dao.SysRecharge.Columns().Money))))
			_, err = tx.Model(dao.SysUser.Table()).WherePri(v.GMap().Get(dao.SysRecharge.Columns().UserId)).Data(g.Map{
				dao.SysUser.Columns().Balance: newbalance,
			}).Update()
			if err != nil {
				return utils_error.Err(response.DB_SAVE_ERROR)
			}
			entity[dao.SysBalance.Columns().Before] = newbalance

			_, err = tx.Model(dao.SysBalance.Table()).
				Data(entity).Insert()
			if err != nil {
				return utils_error.Err(response.DB_SAVE_ERROR)
			}
		}
	}

	_, err = tx.Model(dao.SysRecharge.Table()).
		WhereIn(dao.SysRecharge.Columns().Id, ids).
		Delete()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
