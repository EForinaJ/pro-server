package v1

import (
	dao_permission "server/app/admin/dao/permission"
	dto_permission "server/app/admin/dto/permission"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllReq struct {
	g.Meta `path:"/permission/all" method:"get" tags:"权限" summary:"获取所有的权限"`
}
type GetAllRes struct {
	List []*dao_permission.List `json:"list" dc:"权限列表"`
}

type CreateReq struct {
	g.Meta `path:"/permission/create"  method:"post" tags:"权限" summary:"创建权限"`
	*dto_permission.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/permission/edit" method:"get" tags:"角色" summary:"获取角色编辑信息"`
	Id     int64 `p:"id" v:"required#请输入id" dc:"角色id"`
}
type GetEditRes struct {
	*dao_permission.Edit
}

type EditReq struct {
	g.Meta `path:"/permission/edit"  method:"post" tags:"权限" summary:"菜单权限"`
	*dto_permission.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/permission/delete"  method:"post" tags:"权限" summary:"权限删除"`
	Ids    []int64 `p:"ids" v:"required|array#请输入菜单id|删除id是数组" dc:"权限id"`
}
type DeleteRes struct{}
