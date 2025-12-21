package order

import (
	"context"
	dto_order "server/app/admin/dto/order"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/do"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Distribute implements service.IOrder.
func (s *sOrder) Distribute(ctx context.Context, req *dto_order.Distribute) (err error) {
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

	var entity do.SysOrderWitkey
	entity.WitkeyId = req.WitkeyId
	entity.OrderId = req.Id
	_, err = tx.Model(dao.SysOrderWitkey.Table()).
		Data(&entity).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	//  添加接单日志
	_, err = tx.Model(dao.SysOrderWitkeyAcceptLog.Table()).Data(g.Map{
		dao.SysOrderWitkeyAcceptLog.Columns().OrderId:    req.Id,
		dao.SysOrderWitkeyAcceptLog.Columns().CreateTime: gtime.Now(),
		dao.SysOrderWitkeyAcceptLog.Columns().WitkeyId:   req.WitkeyId,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	witkeyUser, err := tx.Model(dao.SysWitkey.Table()).
		WherePri(req.WitkeyId).Value(dao.SysWitkey.Columns().UserId)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	user, err := tx.Model(dao.SysUser.Table()).
		WherePri(witkeyUser.Int64()).Value(dao.SysUser.Columns().Name)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	//  添加订单日志
	_, err = tx.Model(dao.SysOrderLog.Table()).Data(g.Map{
		dao.SysOrderLog.Columns().OrderId:    req.Id,
		dao.SysOrderLog.Columns().CreateTime: gtime.Now(),
		dao.SysOrderLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypeDistribute,
		dao.SysOrderLog.Columns().Content:    "派发订单给" + user.String(),
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	return
}
