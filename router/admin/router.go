package router

import (
	"server/app/admin/controller/account"
	"server/app/admin/controller/category"
	"server/app/admin/controller/distribute"
	"server/app/admin/controller/level"
	"server/app/admin/controller/settlement"
	"server/app/admin/controller/witkey"

	"server/app/admin/controller/aftersales"
	"server/app/admin/controller/auth"
	"server/app/admin/controller/bill"
	"server/app/admin/controller/game"
	"server/app/admin/controller/recharge"

	"server/app/admin/controller/dict_data"
	"server/app/admin/controller/dict_type"
	"server/app/admin/controller/manage"
	"server/app/admin/controller/media"
	"server/app/admin/controller/menu"
	"server/app/admin/controller/order"
	"server/app/admin/controller/permission"
	"server/app/admin/controller/product"
	"server/app/admin/controller/role"
	"server/app/admin/controller/site"
	"server/app/admin/controller/system"
	"server/app/admin/controller/title"
	"server/app/admin/controller/upload"
	"server/app/admin/controller/user"
	"server/app/admin/controller/withdraw"
	"server/app/common/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func LoadRouter(s *ghttp.Server) {
	s.Group("/api/v1/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			auth.NewV1(),
			// common.NewV1(),
			site.NewV1(),
		).Middleware(middleware.Response)
		group.Bind(
			account.NewV1(),
		).Middleware(middleware.AdminAuth).Middleware(middleware.Response)
		group.Middleware(middleware.AdminAuth)
		group.Middleware(middleware.Casbin)
		group.Middleware(middleware.Response)
		group.Bind(
			menu.NewV1(),
			manage.NewV1(),
			role.NewV1(),
			permission.NewV1(),
			dict_type.NewV1(),
			dict_data.NewV1(),
			media.NewV1(),
			upload.NewV1(),

			user.NewV1(),
			witkey.NewV1(),
			level.NewV1(),
			game.NewV1(),
			category.NewV1(),
			title.NewV1(),

			product.NewV1(),
			order.NewV1(),
			distribute.NewV1(),
			settlement.NewV1(),
			aftersales.NewV1(),
			recharge.NewV1(),

			bill.NewV1(),
			withdraw.NewV1(),
			system.NewV1(),
		// grade.NewV1(),
		// vip.NewV1(),

		)
	})
}
