package settlement

import (
	"context"
	dto_settlement "server/app/admin/dto/settlement"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/util/gconv"
)

// CheckApply implements service.ISettlement.
func (s *sSettlement) CheckApply(ctx context.Context, req *dto_settlement.Apply) (err error) {
	obj, err := dao.SysSettlement.Ctx(ctx).Where(dao.SysSettlement.Columns().Id, req.Id).
		Fields(dao.SysSettlement.Columns().OrderId, dao.SysSettlement.Columns().Status).
		One()
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}

	orderStatus, err := dao.SysOrder.Ctx(ctx).
		Where(dao.SysOrder.Columns().Id, obj.GMap().Get(dao.SysSettlement.Columns().OrderId)).
		Value(dao.SysOrder.Columns().Status)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	if orderStatus.Int() != consts.OrderStatusComplete {
		return utils_error.ErrMessage(response.FAILD, "订单未完成，无法审核结算")
	}

	if gconv.Int(obj.GMap().Get(dao.SysSettlement.Columns().Status)) != consts.StatusApply {
		return utils_error.ErrMessage(response.FAILD, "该报单结算已审核")
	}

	return
}
