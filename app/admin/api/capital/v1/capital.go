package v1

import (
	dao_capital "server/app/admin/dao/capital"
	dto_capital "server/app/admin/dto/capital"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/capital/list" method:"get" tags:"资金" summary:"资金记录列表"`
	*dto_capital.Query
}
type GetListRes struct {
	Total int                 `json:"total" dc:"总数"`
	List  []*dao_capital.List `json:"list" dc:"资金记录列表"`
}

type DeleteReq struct {
	g.Meta `path:"/capital/delete" method:"post" tags:"资金" summary:"删除资金记录"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
