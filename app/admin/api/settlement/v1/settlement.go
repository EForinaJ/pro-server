package v1

import (
	dao_settlement "server/app/admin/dao/settlement"
	dto_settlement "server/app/admin/dto/settlement"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/settlement/list" method:"get" tags:"结算" summary:"结算列表"`
	*dto_settlement.Query
}
type GetListRes struct {
	Total int                    `json:"total" dc:"总数"`
	List  []*dao_settlement.List `json:"list" dc:"结算列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/settlement/detail" method:"get" tags:"结算" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_settlement.Detail
}

type ApplyReq struct {
	g.Meta `path:"/settlement/apply" method:"post" tags:"结算" summary:"审核结算"`
	*dto_settlement.Apply
}
type ApplyRes struct{}
