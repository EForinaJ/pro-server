package router

import (
	"server/app/common/middleware"
	"server/app/frontend/controller/account"
	"server/app/frontend/controller/bill"
	"server/app/frontend/controller/experience"
	"server/app/frontend/controller/level"
	"server/app/frontend/controller/order"
	"server/app/frontend/controller/recharge"
	"server/app/frontend/controller/upload"

	"server/app/frontend/controller/auth"
	"server/app/frontend/controller/document"
	"server/app/frontend/controller/site"

	"github.com/gogf/gf/v2/net/ghttp"
)

func LoadRouter(s *ghttp.Server) {
	s.Group("/api/v1/frontend", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Middleware(ghttp.MiddlewareCORS)
		group.Middleware(middleware.FixedWindowMiddleware)
		group.Bind(
			auth.NewV1(),
			site.NewV1(),
			document.NewV1(),
			recharge.NewV1().GetNotify,
			order.NewV1().GetNotify,
		).Middleware(middleware.Response)

		// 读取操作
		group.Bind(
			account.NewV1().GetDetail,
			recharge.NewV1().GetStatus,
			order.NewV1().GetStatus,

			order.NewV1().GetList,
			order.NewV1().GetDetail,

			bill.NewV1().GetList,
			level.NewV1().GetAll,
			experience.NewV1().GetList,

			account.NewV1().GetCardList,
			account.NewV1().GetWithdrawList,
		).Middleware(middleware.FrontendAuth).Middleware(middleware.Response)
		// 写入操作
		group.Bind(
			upload.NewV1(),
			account.NewV1().WxProgramBindPhone,
			account.NewV1().Edit,

			recharge.NewV1().WechatMiniProgram, // 小程序充值
			order.NewV1().Cancel,
			order.NewV1().Delete,
			order.NewV1().BalancePay,
			order.NewV1().WechatMiniProgramPay,

			// account.NewV1().ChangePassword,
			// account.NewV1().SignIn,
			// comment.NewV1().Create,
			// account.NewV1().CreateCard,
			// account.NewV1().DeleteCard,
			// account.NewV1().Withdraw,
		).Middleware(middleware.FixedWindowActionMiddleware).
			Middleware(middleware.FrontendAuth).
			Middleware(middleware.Response)
	})
}
