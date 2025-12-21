package v1

import (
	dao_level "server/app/admin/dao/level"
	dto_level "server/app/admin/dto/level"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllReq struct {
	g.Meta `path:"/level/all" method:"get" tags:"等级" summary:"所以等级"`
}
type GetAllRes struct {
	List []*dao_level.List `json:"list" dc:"等级列表"`
}

type GetListReq struct {
	g.Meta `path:"/level/list" method:"get" tags:"等级" summary:"等级列表"`
	*dto_level.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"总数"`
	List  []*dao_level.List `json:"list" dc:"等级列表"`
}

type CreateReq struct {
	g.Meta `path:"/level/create" method:"post" tags:"等级" summary:"创建等级"`
	*dto_level.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/level/edit" method:"get" tags:"等级" summary:"获取编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetEditRes struct {
	*dao_level.Edit
}

type EditReq struct {
	g.Meta `path:"/level/edit" method:"post" tags:"等级" summary:"等级编辑"`
	*dto_level.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/level/delete" method:"post" tags:"等级" summary:"删除等级"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
