package v1

import (
	dao_account "server/app/admin/dao/account"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDetailReq struct {
	g.Meta `path:"/account" method:"get" tags:"账户" summary:"获取登录账户信息"`
}
type GetDetailRes struct {
	*dao_account.Detail
}

type GetMenuReq struct {
	g.Meta `path:"/account/menus" method:"get" tags:"账户" summary:"获取登录账户路由"`
}
type GetMenuRes struct {
	List []*dao_account.Menu `json:"list" dc:"菜单列表"`
}

type GetDashboardReq struct {
	g.Meta `path:"/account/dashboard" method:"get" tags:"账户" summary:"获取登录账户数据"`
}
type GetDashboardRes struct {
	*dao_account.Dashboard
}
