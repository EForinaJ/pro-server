package order

import (
	"context"
	dao_order "server/app/admin/dao/order"
	dto_order "server/app/admin/dto/order"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/util/gconv"
)

type sOrder struct{}

// GetLogs implements service.IOrder.
func (s *sOrder) GetLogs(ctx context.Context, req *dto_order.Log) (total int, res []*dao_order.Log, err error) {
	m := dao.SysOrderLog.Ctx(ctx).
		Page(req.Page, req.Limit).
		Where(dao.SysOrderLog.Columns().OrderId, req.Id).
		OrderDesc(dao.SysOrderLog.Columns().CreateTime)

	total, err = m.Count()
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}
	var list []*entity.SysOrderLog
	err = m.Scan(&list)
	if err != nil {
		return 0, nil, utils_error.Err(response.DB_READ_ERROR)
	}

	res = make([]*dao_order.Log, len(list))
	for i, v := range list {
		var entity *dao_order.Log
		err = gconv.Scan(v, &entity)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}

		//  操作人
		user, err := dao.SysManage.Ctx(ctx).
			Where(dao.SysManage.Columns().Id, v.ManageId).
			Value(dao.SysManage.Columns().Name)
		if err != nil {
			return 0, nil, utils_error.Err(response.DB_READ_ERROR)
		}
		entity.Name = gconv.String(user)

		res[i] = entity
	}
	return
}

func init() {
	service.RegisterOrder(&sOrder{})
}
