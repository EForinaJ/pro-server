package order

import (
	"context"
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

// BalancePay implements service.IOrder.
func (s *sOrder) BalancePay(ctx context.Context, id int64) (err error) {
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

	order, err := tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, id).
		Where(dao.SysOrder.Columns().UserId, ctx.Value("userId")).
		Fields(dao.SysOrder.Columns().ActualAmount, dao.SysOrder.Columns().Code).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	actualAmount := decimal.NewFromFloat(gconv.Float64(order.GMap().Get(dao.SysOrder.Columns().ActualAmount)))

	userBalance, err := tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Value(dao.SysUser.Columns().Balance)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	balance := decimal.NewFromFloat(userBalance.Float64())
	newBalance := balance.Sub(actualAmount)

	_, err = tx.Model(dao.SysUser.Table()).
		Where(dao.SysUser.Columns().Id, ctx.Value("userId")).
		Data(g.Map{
			dao.SysUser.Columns().Balance: newBalance,
		}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	_, err = tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, id).
		Data(g.Map{
			dao.SysOrder.Columns().PayStatus: consts.PayStatusSuccess,
			dao.SysOrder.Columns().PayTime:   gtime.Now(),
			dao.SysOrder.Columns().PayMode:   consts.PayModeBalance,
			// dao.SysOrder.Columns().PayCode:   utils_snow.GetCode(ctx, "YE"),
			dao.SysOrder.Columns().PayAmount: actualAmount,
			dao.SysOrder.Columns().Status:    consts.OrderStatusPendingOrder,
		}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	billEntity := g.Map{
		dao.SysUserBill.Columns().UserId:     ctx.Value("userId"),
		dao.SysUserBill.Columns().RelatedId:  id,
		dao.SysUserBill.Columns().Code:       utils_snow.GetCode(ctx, consts.BL),
		dao.SysUserBill.Columns().Type:       consts.BillTypeOrder,
		dao.SysUserBill.Columns().Amount:     actualAmount,
		dao.SysUserBill.Columns().Mode:       consts.Sub,
		dao.SysUserBill.Columns().CreateTime: gtime.Now(),
	}
	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(billEntity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = g.Redis().Del(ctx, consts.FrontendAccout+gconv.String(ctx.Value("userId")))
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
