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
	"github.com/gogf/gf/v2/util/gconv"
)

// Create implements service.IOrder.
func (s *sOrder) Create(ctx context.Context, req *dto_order.Create) (err error) {
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

	var entity *do.SysOrder
	err = gconv.Scan(req, &entity)
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}
	entity.CreateTime = gtime.Now()
	entity.UpdateTime = gtime.Now()
	entity.Status = consts.OrderStatusPendingPayment

	// 获取订单金额
	// goods, err := dao.SysGoods.Ctx(ctx).
	// 	Fields(
	// 		dao.SysGoods.Columns().Unit,
	// 		dao.SysGoods.Columns().Price).
	// 	WherePri(req.GoodsId).
	// 	One()
	// if err != nil {
	// 	return utils_error.Err(response.DB_SAVE_ERROR)
	// }
	// money := gconv.Float64(goods.GMap().Get(dao.SysGoods.Columns().Price))
	// price := decimal.NewFromFloat(money).Mul(decimal.NewFromFloat(gconv.Float64(req.Number)))
	// entity.Money = price
	// entity.PayStatus = consts.PayStatusPending
	// entity.Unit = goods.GMap().Get(dao.SysGoods.Columns().Unit)
	// entity.Code = utils_snow.GenCode(ctx, req.UserId)
	// manageId := ctx.Value("userId")
	// entity.ManageId = manageId

	_, err = tx.Model(dao.SysOrder.Table()).Data(entity).Insert()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
