package order

import (
	"context"
	dto_order "server/app/admin/dto/order"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChangeWitkey implements service.IOrder.
func (s *sOrder) ChangeWitkey(ctx context.Context, req *dto_order.ChangeWitkey) (err error) {
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

	_, err = tx.Model(dao.SysOrderWitkey.Table()).
		Data(g.Map{
			dao.SysOrderWitkey.Columns().IsReplaced: consts.Not,
			dao.SysOrderWitkey.Columns().WitkeyId:   req.NewId,
			dao.SysOrderWitkey.Columns().OrderId:    req.Id,
			dao.SysOrderWitkey.Columns().CreateTime: gtime.Now(),
		}).
		Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = tx.Model(dao.SysOrderWitkey.Table()).
		Where(dao.SysOrderWitkey.Columns().WitkeyId, req.OldId).
		Where(dao.SysOrderWitkey.Columns().OrderId, req.Id).
		Data(g.Map{
			dao.SysOrderWitkey.Columns().IsReplaced: consts.Yes,
			dao.SysOrderWitkey.Columns().Reason:     req.Reason,
		}).
		Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	_, err = tx.Model(dao.SysOrderLog.Table()).Data(g.Map{
		dao.SysOrderLog.Columns().OrderId:    req.Id,
		dao.SysOrderLog.Columns().CreateTime: gtime.Now(),
		dao.SysOrderLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysOrderLog.Columns().Content:    "变更原因：" + req.Reason,
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypeChangeWitkey,
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
