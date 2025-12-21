package account

import (
	"context"
	dao_account "server/app/admin/dao/account"
	"server/app/common/consts"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	"server/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetDashboard implements service.IAccount.
func (s *sAccount) GetDashboard(ctx context.Context) (res *dao_account.Dashboard, err error) {
	options, err := g.Redis().Get(ctx, consts.AdminDashboard+gconv.String(ctx.Value("userId")))
	if err != nil {
		return nil, utils_error.Err(response.CACHE_READ_ERROR)
	}
	if !options.IsEmpty() {
		err = options.Scan(&res)
		if err != nil {
			return nil, utils_error.Err(response.CACHE_READ_ERROR)
		}
		return
	}

	// 获取今日0点和明日0点
	startOfToday := gtime.Now().StartOfDay()
	startOfTomorrow := startOfToday.AddDate(0, 0, 1)

	// 获取用户数量
	userAllCount, err := dao.SysUser.Ctx(ctx).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	userTodayCount, err := dao.SysUser.Ctx(ctx).
		Where("create_time >= ?", startOfToday).
		Where("create_time < ?", startOfTomorrow).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	userCount := g.Map{
		"today": userTodayCount,
		"total": userAllCount,
	}

	// 获取订单数量
	orderAllCount, err := dao.SysOrder.Ctx(ctx).Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	orderTodayCount, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday).
		Where("create_time < ?", startOfTomorrow).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	orderCount := g.Map{
		"today": orderTodayCount,
		"total": orderAllCount,
	}

	// 获取预存金额
	rechargeAllMoney, err := dao.SysRecharge.Ctx(ctx).Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	rechargeTodayMoney, err := dao.SysRecharge.Ctx(ctx).
		Where("create_time >= ?", startOfToday).
		Where("create_time < ?", startOfTomorrow).
		Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	rechargeMoney := g.Map{
		"today": rechargeTodayMoney,
		"total": rechargeAllMoney,
	}

	// 获取订单数量
	witkeyAllCount, err := dao.SysWitkey.Ctx(ctx).Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	witkeyTodayCount, err := dao.SysWitkey.Ctx(ctx).
		Where("create_time >= ?", startOfToday).
		Where("create_time < ?", startOfTomorrow).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	witkeyCount := g.Map{
		"today": witkeyTodayCount,
		"total": witkeyAllCount,
	}
	var dashboard dao_account.Dashboard

	dashboard.CountCard = g.Map{
		"witkeyCount":   witkeyCount,
		"rechargeMoney": rechargeMoney,
		"orderCount":    orderCount,
		"userCount":     userCount,
	}

	// 待支付订单
	peddingPayment, err := dao.SysOrder.
		Ctx(ctx).Where(dao.SysOrder.Columns().Status, consts.OrderStatusPendingPayment).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	// 待服务
	peddingService, err := dao.SysOrder.
		Ctx(ctx).Where(dao.SysOrder.Columns().Status, consts.OrderStatusPendingOrder).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	// 服务中
	inProgress, err := dao.SysOrder.
		Ctx(ctx).Where(dao.SysOrder.Columns().Status, consts.OrderStatusInProgress).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	// 已完成
	complete, err := dao.SysOrder.
		Ctx(ctx).Where(dao.SysOrder.Columns().Status, consts.OrderStatusComplete).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	// 已取消
	cancel, err := dao.SysOrder.
		Ctx(ctx).Where(dao.SysOrder.Columns().Status, consts.OrderStatusCancel).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	// 已取消
	aftersales, err := dao.SysOrder.
		Ctx(ctx).Where(dao.SysOrder.Columns().Status, consts.OrderStatusRefund).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	dashboard.OrderStatistics = g.Map{
		"peddingPayment": peddingPayment,
		"peddingService": peddingService,
		"inProgress":     inProgress,
		"complete":       complete,
		"cancel":         cancel,
		"aftersales":     aftersales,
	}

	// 未消费过用户数量
	unconsumed, err := dao.SysUser.Ctx(ctx).As("u").
		LeftJoin("sys_order as o", "u.id = o.user_id").Where("o.id IS NULL").Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	// 多次消费的用户数量
	multipleSales, err := dao.SysOrder.Ctx(ctx).Fields("user_id").
		Group("user_id").Having("COUNT(*) > 1").Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	// 留存消费用户数量
	subQuery := g.Model("sys_recharge o").Where("o.user_id = u.id")
	recharge, err := dao.SysUser.Ctx(ctx).As("u").
		WhereExists(subQuery).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	dashboard.UserConsumptionStatistics = g.Map{
		"unconsumed":    unconsumed,
		"multipleSales": multipleSales,
		"recharge":      recharge,
	}

	oneOrderCount, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday).
		Where("create_time < ?", startOfTomorrow).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	oneOrderMoney, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday).
		Where("create_time < ?", startOfTomorrow).
		Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	subOneOrderCount, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -1)).
		Where("create_time < ?", startOfToday).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	subOneOrderMoney, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -1)).
		Where("create_time < ?", startOfToday).
		Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	subTwoOrderCount, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -2)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -1)).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	subTwoOrderMoney, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -2)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -1)).
		Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	subThreeOrderCount, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -3)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -2)).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	subThreeOrderMoney, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -3)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -2)).
		Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	subFourOrderCount, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -4)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -3)).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	subFourOrderMoney, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -4)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -3)).
		Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	subFiveOrderCount, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -5)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -4)).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	subFiveOrderMoney, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -5)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -4)).
		Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	subSixOrderCount, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -6)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -5)).
		Count()
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}
	subSixOrderMoney, err := dao.SysOrder.Ctx(ctx).
		Where("create_time >= ?", startOfToday.AddDate(0, 0, -6)).
		Where("create_time < ?", startOfToday.AddDate(0, 0, -5)).
		Sum("money")
	if err != nil {
		return nil, utils_error.Err(response.DB_READ_ERROR)
	}

	lineOrderStatistics := g.Map{
		"date": []string{
			startOfToday.AddDate(0, 0, -6).Format("Y-m-d"),
			startOfToday.AddDate(0, 0, -5).Format("Y-m-d"),
			startOfToday.AddDate(0, 0, -4).Format("Y-m-d"),
			startOfToday.AddDate(0, 0, -3).Format("Y-m-d"),
			startOfToday.AddDate(0, 0, -2).Format("Y-m-d"),
			startOfToday.AddDate(0, 0, -1).Format("Y-m-d"),
			startOfToday.Format("Y-m-d"),
		},
		"orderCount": []int{
			subSixOrderCount,
			subFiveOrderCount,
			subFourOrderCount,
			subThreeOrderCount,
			subTwoOrderCount,
			subOneOrderCount,
			oneOrderCount,
		},
		"orderMoney": []float64{
			subSixOrderMoney,
			subFiveOrderMoney,
			subFourOrderMoney,
			subThreeOrderMoney,
			subTwoOrderMoney,
			subOneOrderMoney,
			oneOrderMoney,
		},
	}

	dashboard.LineOrderStatistics = lineOrderStatistics

	// goodsList, err := dao.SysGoods.Ctx(ctx).
	// 	Fields(dao.SysGoods.Columns().Id,
	// 		dao.SysGoods.Columns().Name,
	// 		dao.SysGoods.Columns().SalesCount,
	// 		dao.SysGoods.Columns().Price,
	// 		dao.SysGoods.Columns().Cover).
	// 	OrderDesc(dao.SysGoods.Columns().SalesCount).Page(1, 5).All()
	// if err != nil {
	// 	return nil, utils_error.Err(response.DB_READ_ERROR)
	// }
	// hotGoodsListres := make([]g.Map, len(goodsList))
	// for i, v := range goodsList {
	// 	goods := g.Map{
	// 		"id":         v.GMap().Get(dao.SysGoods.Columns().Id),
	// 		"name":       v.GMap().Get(dao.SysGoods.Columns().Name),
	// 		"salesCount": v.GMap().Get(dao.SysGoods.Columns().SalesCount),
	// 		"price":      v.GMap().Get(dao.SysGoods.Columns().Price),
	// 		"cover":      v.GMap().Get(dao.SysGoods.Columns().Cover),
	// 	}
	// 	// 获取威客人数
	// 	hotGoodsListres[i] = goods
	// }
	// dashboard.HotGoods = hotGoodsListres

	err = g.Redis().SetEX(ctx, consts.AdminDashboard+gconv.String(ctx.Value("userId")), dashboard, 600)
	if err != nil {
		return nil, utils_error.Err(response.CACHE_SAVE_ERROR)
	}
	return &dashboard, nil
}
