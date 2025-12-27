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
)

// Paid implements service.IOrder.
func (s *sOrder) Paid(ctx context.Context, id int64) (err error) {
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
		Fields(
			dao.SysOrder.Columns().ActualAmount,
			dao.SysOrder.Columns().UserId,
			dao.SysOrder.Columns().Code,
		).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	_, err = tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, id).
		Data(g.Map{
			dao.SysOrder.Columns().Status:    consts.OrderStatusPendingOrder,
			dao.SysOrder.Columns().PayStatus: consts.PayStatusSuccess,
			dao.SysOrder.Columns().PayTime:   gtime.Now(),
			dao.SysOrder.Columns().PayMode:   consts.PayModePersonalTransfer,
			dao.SysOrder.Columns().PayAmount: order.GMap().Get(dao.SysOrder.Columns().ActualAmount),
		}).
		Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(g.Map{
			dao.SysUserBill.Columns().UserId:     order.GMap().Get(dao.SysOrder.Columns().UserId),
			dao.SysUserBill.Columns().RelatedId:  id,
			dao.SysUserBill.Columns().Code:       utils_snow.GetCode(ctx, consts.BL),
			dao.SysUserBill.Columns().Type:       consts.BillTypeOrder,
			dao.SysUserBill.Columns().Amount:     order.GMap().Get(dao.SysOrder.Columns().ActualAmount),
			dao.SysUserBill.Columns().Mode:       consts.Sub,
			dao.SysUserBill.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	//  添加支付日志
	_, err = tx.Model(dao.SysCapital.Table()).Data(g.Map{
		dao.SysCapital.Columns().CreateTime: gtime.Now(),
		dao.SysCapital.Columns().Code:       utils_snow.GetCode(ctx, consts.PM),
		dao.SysCapital.Columns().Related:    order.GMap().Get(dao.SysOrder.Columns().Code),
		dao.SysCapital.Columns().Amount:     order.GMap().Get(dao.SysOrder.Columns().ActualAmount),
		dao.SysCapital.Columns().Type:       consts.CapitalPaymentOrder,
		dao.SysCapital.Columns().Mode:       consts.PayModePersonalTransfer,
		dao.SysCapital.Columns().UserId:     order.GMap().Get(dao.SysOrder.Columns().UserId),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
