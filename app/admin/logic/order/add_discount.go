package order

import (
	"context"
	dto_order "server/app/admin/dto/order"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/shopspring/decimal"
)

// AddDiscount implements service.IOrder.
func (s *sOrder) AddDiscount(ctx context.Context, req *dto_order.AddDiscount) (err error) {
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

	orderTotal, err := dao.SysOrder.Ctx(ctx).Where(dao.SysOrder.Columns().Id, req.Id).
		Value(dao.SysOrder.Columns().TotalAmount)
	if err != nil {
		return utils_error.Err(response.DB_READ_ERROR)
	}
	discountAmout := decimal.NewFromFloat(req.Amount)
	actualAmount := decimal.NewFromFloat(orderTotal.Float64()).Sub(discountAmout)

	_, err = tx.Model(dao.SysOrder.Table()).WherePri(req.Id).Data(g.Map{
		dao.SysOrder.Columns().DiscountAmount: req.Amount,
		dao.SysOrder.Columns().ActualAmount:   actualAmount,
	}).Update()
	if err != nil {
		return utils_error.Err(response.DB_SAVE_ERROR)
	}

	return
}
