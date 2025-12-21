package v1

import (
	dao_menu "server/app/admin/dao/menu"
	dao_permission "server/app/admin/dao/permission"
	dao_role "server/app/admin/dao/role"
	dto_role "server/app/admin/dto/role"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllMenuReq struct {
	g.Meta `path:"/role/all/menu" method:"get" tags:"角色" summary:"获取角色所有菜单"`
}
type GetAllMenuRes struct {
	List []*dao_menu.List `json:"list" dc:"所有菜单"`
}

type GetAllPermissionReq struct {
	g.Meta `path:"/role/all/permission" method:"get" tags:"角色" summary:"获取角色所有权限"`
}
type GetAllPermissionRes struct {
	List []*dao_permission.List `json:"list" dc:"所有权限"`
}

type EditPermissionsReq struct {
	g.Meta      `path:"/role/edit/permissions" method:"post" tags:"角色" summary:"修改角色所拥有接口权限"`
	Id          int64   `p:"id" v:"required|integer|min:1#请输入角色id|角色id类型必须是整型|角色id最小为1" dc:"角色id"`
	Permissions []int64 `p:"permissions" v:"required|array#请设置接口权限|接口权限必须是一个数组" dc:"接口权限Id数组"`
}
type EditPermissionsRes struct{}

type GetPermissionsReq struct {
	g.Meta `path:"/role/permissions" method:"get" tags:"角色" summary:"获取角色所拥有的接口权限"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"角色id"`
}
type GetPermissionsRes struct {
	Permissions []int64 `json:"permissiones" dc:"角色拥有的接口权限"`
}

type EditMenusReq struct {
	g.Meta `path:"/role/edit/menus" method:"post" tags:"角色" summary:"修改角色所拥有菜单"`
	Id     int64   `p:"id" v:"required|integer|min:1#请输入角色id|角色id类型必须是整型|角色id最小为1" dc:"角色id"`
	Menus  []int64 `p:"menus" v:"required|array#请设置菜单|菜单必须是一个数组" dc:"菜单Id数组"`
}
type EditMenusRes struct{}

type GetMenusReq struct {
	g.Meta `path:"/role/menus" method:"get" tags:"角色" summary:"获取角色所拥有的菜单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"角色id"`
}
type GetMenusRes struct {
	Menus []int64 `json:"menus" dc:"角色拥有的菜单"`
}

type GetListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色" summary:"角色列表"`
	*dto_role.Query
}
type GetListRes struct {
	Total int              `json:"total" dc:"总数"`
	List  []*dao_role.List `json:"list" dc:"角色列表"`
}

type CreateReq struct {
	g.Meta `path:"/role/create" method:"post" tags:"角色" summary:"创建角色"`
	*dto_role.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/role/edit" method:"get" tags:"角色" summary:"获取角色编辑信息"`
	Id     int64 `p:"id" v:"required#请输入id" dc:"角色id"`
}
type GetEditRes struct {
	*dao_role.Edit
}

type EditReq struct {
	g.Meta `path:"/role/edit" method:"post" tags:"角色" summary:"修改角色"`
	*dto_role.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/role/delete" method:"post" tags:"角色" summary:"删除角色"`
	Ids    []int64 `p:"ids" v:"required|array#请输入id|删除id是数组" dc:"角色id"`
}
type DeleteRes struct{}
