package v1

import (
	dao_dict_data "server/app/admin/dao/dict_data"
	dto_dict_data "server/app/admin/dto/dict_data"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllReq struct {
	g.Meta `path:"/dict/data/all" method:"get" tags:"字典" summary:"字典数据列表"`
	Code   string `p:"code" v:"required#请输入字典码" dc:"字典码"`
}
type GetAllRes struct {
	List []*dao_dict_data.List `json:"list" dc:"字典数据列表"`
}

type GetListReq struct {
	g.Meta `path:"/dict/data/list" method:"get" tags:"字典" summary:"字典数据列表"`
	*dto_dict_data.Query
}
type GetListRes struct {
	Total int                   `json:"total" dc:"总数"`
	List  []*dao_dict_data.List `json:"list" dc:"字典数据列表"`
}

type CreateReq struct {
	g.Meta `path:"/dict/data/create" method:"post" tags:"字典" summary:"创建字典数据"`
	*dto_dict_data.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/dict/data/edit" method:"get" tags:"字典" summary:"获取字典数据编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetEditRes struct {
	*dao_dict_data.Edit
}

type EditReq struct {
	g.Meta `path:"/dict/data/edit" method:"post" tags:"字典" summary:"字典数据编辑"`
	*dto_dict_data.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/dict/data/delete" method:"post" tags:"字典" summary:"删除字典数据"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
