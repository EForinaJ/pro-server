package order

import (
	"context"
	dto_order "server/app/admin/dto/order"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	utils_snow "server/app/common/utils/snow"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Refund implements service.IOrder.
func (s *sOrder) Refund(ctx context.Context, req *dto_order.Refund) (err error) {
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
	order, err := tx.Model(dao.SysOrder.Table()).WherePri(req.Id).
		Fields(dao.SysOrder.Columns().PayMode,
			dao.SysOrder.Columns().UserId,
			dao.SysOrder.Columns().Code).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	_, err = tx.Model(dao.SysOrder.Table()).WherePri(req.Id).Data(g.Map{
		dao.SysOrder.Columns().PayStatus: consts.PayStatusRefund,
		dao.SysOrder.Columns().Status:    consts.OrderStatusRefund,
	}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = tx.Model(dao.SysAftersales.Table()).
		Data(g.Map{
			dao.SysAftersales.Columns().OrderId:  req.Id,
			dao.SysAftersales.Columns().Code:     utils_snow.GetCode(ctx, consts.AS),
			dao.SysAftersales.Columns().Amount:   req.Money,
			dao.SysAftersales.Columns().Type:     req.Type,
			dao.SysAftersales.Columns().ManageId: ctx.Value("userId"),
			dao.SysAftersales.Columns().Reason:   req.Reason,
			dao.SysAftersales.Columns().Status:   consts.StatusSuccess,
			dao.SysUserBill.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	_, err = tx.Model(dao.SysRefund.Table()).
		Data(g.Map{
			dao.SysRefund.Columns().OrderId:      req.Id,
			dao.SysRefund.Columns().Mode:         consts.PayModePersonalTransfer,
			dao.SysRefund.Columns().Money:        req.Money,
			dao.SysRefund.Columns().ManageId:     ctx.Value("userId"),
			dao.SysUserBill.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	orderUserId := gconv.Int64(order.GMap().Get(dao.SysOrder.Columns().UserId))
	_, err = tx.Model(dao.SysUserBill.Table()).
		Data(g.Map{
			dao.SysUserBill.Columns().UserId:     orderUserId,
			dao.SysUserBill.Columns().RelatedId:  req.Id,
			dao.SysUserBill.Columns().Code:       utils_snow.GetCode(ctx, consts.BL),
			dao.SysUserBill.Columns().Type:       consts.BillTypeRefund,
			dao.SysUserBill.Columns().Money:      req.Money,
			dao.SysUserBill.Columns().Mode:       consts.Add,
			dao.SysUserBill.Columns().CreateTime: gtime.Now(),
		}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
