package v1

import (
	dao_manage "server/app/admin/dao/manage"
	dao_role "server/app/admin/dao/role"
	dto_manage "server/app/admin/dto/manage"

	"github.com/gogf/gf/v2/frame/g"
)

type DeleteReq struct {
	g.Meta `path:"/manage/delete" method:"post" tags:"管理员" summary:"删除管理员"`
	Ids    []int64 `p:"ids" v:"required|array#请输入id|删除id是数组" dc:"管理员id"`
}
type DeleteRes struct{}

type GetAllRoleReq struct {
	g.Meta `path:"/manage/all/role" method:"get" tags:"管理员" summary:"获取管理有的角色"`
}
type GetAllRoleRes struct {
	List []*dao_role.List `json:"list" dc:"获取所有角色列表"`
}

type GetListReq struct {
	g.Meta `path:"/manage/list" method:"get" tags:"管理员" summary:"获取管理员列表"`
	*dto_manage.Query
}
type GetListRes struct {
	Total int                `json:"total" dc:"管理员总数"`
	List  []*dao_manage.List `json:"list" dc:"管理员列表"`
}

type CreateReq struct {
	g.Meta `path:"/manage/create"  method:"post" tags:"管理员" summary:"创建"`
	*dto_manage.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/manage/edit" method:"get" tags:"管理员" summary:"获取管理员编辑信息"`
	Id     int64 `p:"id" v:"required#请输入id" dc:"管理员id"`
}
type GetEditRes struct {
	*dao_manage.Edit
}

type EditReq struct {
	g.Meta `path:"/manage/edit"  method:"post" tags:"管理员" summary:"修改"`
	*dto_manage.Edit
}
type EditRes struct{}
