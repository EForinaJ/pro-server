package order

import (
	"context"
	"server/app/common/consts"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
)

// Complete implements service.IOrder.
func (s *sOrder) Complete(ctx context.Context, id int64) (err error) {
	_, err = dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Id, id).Data(g.Map{
		dao.SysOrder.Columns().Status: consts.OrderStatusComplete,
	}).Update()
	return
}
