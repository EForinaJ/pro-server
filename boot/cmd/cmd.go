package cmd

import (
	"context"
	_ "server/app/admin/logic"
	"server/app/common/library/casbin"
	_ "server/app/frontend/logic"

	admin_router "server/router/admin"
	frontend_router "server/router/frontend"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			casbin.InitEnforcer(ctx)
			s := g.Server()
			s.AddStaticPath("/public", "./public")
			admin_router.LoadRouter(s)
			frontend_router.LoadRouter(s)
			// witkey_router.LoadRouter(s)
			s.Run()
			return nil
		},
	}
)
