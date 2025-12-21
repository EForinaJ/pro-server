package order

import (
	"context"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Cancel implements service.IOrder.
func (s *sOrder) Cancel(ctx context.Context, id int64) (err error) {
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

	_, err = tx.Model(dao.SysOrder.Table()).
		Where(dao.SysOrder.Columns().Id, id).
		Data(g.Map{
			dao.SysOrder.Columns().Status: consts.OrderStatusCancel,
		}).
		Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	//  添加日志
	_, err = tx.Model(dao.SysOrderLog.Table()).Data(g.Map{
		dao.SysOrderLog.Columns().OrderId:    id,
		dao.SysOrderLog.Columns().CreateTime: gtime.Now(),
		dao.SysOrderLog.Columns().ManageId:   ctx.Value("userId"),
		dao.SysOrderLog.Columns().Type:       consts.OrderLogTypeCancel,
		dao.SysOrderLog.Columns().Content:    "关闭订单",
	}).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	return
}
