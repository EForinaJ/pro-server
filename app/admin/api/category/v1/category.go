package v1

import (
	dao_category "server/app/admin/dao/category"
	dto_category "server/app/admin/dto/category"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"分类" summary:"分类列表"`
	*dto_category.Query
}
type GetListRes struct {
	Total int                  `json:"total" dc:"总数"`
	List  []*dao_category.List `json:"list" dc:"分类列表"`
}

type CreateReq struct {
	g.Meta `path:"/category/create" method:"post" tags:"分类" summary:"创建分类"`
	*dto_category.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/category/edit" method:"get" tags:"分类" summary:"获取编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetEditRes struct {
	*dao_category.Edit
}

type EditReq struct {
	g.Meta `path:"/category/edit" method:"post" tags:"分类" summary:"分类编辑"`
	*dto_category.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/category/delete" method:"post" tags:"分类" summary:"删除分类"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
