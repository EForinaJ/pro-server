package v1

import (
	dao_menu "server/app/admin/dao/menu"
	dto_meun "server/app/admin/dto/meun"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllReq struct {
	g.Meta `path:"/menu/all" method:"get" tags:"菜单" summary:"获取所有的菜单"`
}
type GetAllRes struct {
	List []*dao_menu.List `json:"list" dc:"菜单列表"`
}

type CreateReq struct {
	g.Meta `path:"/menu/create"  method:"post" tags:"菜单" summary:"创建菜单"`
	*dto_meun.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/menu/edit" method:"get" tags:"菜单" summary:"获取菜单编辑信息"`
	Id     int64 `p:"id" v:"required#请输入id" dc:"菜单id"`
}
type GetEditRes struct {
	*dao_menu.Edit
}

type EditReq struct {
	g.Meta `path:"/menu/edit"  method:"post" tags:"菜单" summary:"菜单修改"`
	*dto_meun.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/menu/delete"  method:"post" tags:"菜单" summary:"菜单删除"`
	Ids    []int64 `p:"ids" v:"required|array#请输入菜单id|删除id是数组" dc:"菜单id"`
}
type DeleteRes struct{}
