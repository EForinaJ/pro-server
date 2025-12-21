package v1

import (
	dao_dict_type "server/app/admin/dao/dict_type"
	dto_dict_type "server/app/admin/dto/dict_type"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/dict/type/list" method:"get" tags:"字典" summary:"字典列表"`
	*dto_dict_type.Query
}
type GetListRes struct {
	Total int                   `json:"total" dc:"总数"`
	List  []*dao_dict_type.List `json:"list" dc:"字典列表"`
}

type CreateReq struct {
	g.Meta `path:"/dict/type/create" method:"post" tags:"字典" summary:"创建字典"`
	*dto_dict_type.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/dict/type/edit" method:"get" tags:"字典" summary:"获取编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetEditRes struct {
	*dao_dict_type.Edit
}

type EditReq struct {
	g.Meta `path:"/dict/type/edit" method:"post" tags:"字典" summary:"字典编辑"`
	*dto_dict_type.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/dict/type/delete" method:"post" tags:"字典" summary:"删除字典"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
