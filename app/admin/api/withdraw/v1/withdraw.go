package v1

import (
	dao_withdraw "server/app/admin/dao/withdraw"
	dto_withdraw "server/app/admin/dto/withdraw"

	"github.com/gogf/gf/v2/frame/g"
)

type GetInfoReq struct {
	g.Meta `path:"/withdraw/detail" method:"get" tags:"提现申请" summary:"提现详情"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetInfoRes struct {
	*dao_withdraw.Detail
}

type GetListReq struct {
	g.Meta `path:"/withdraw/list" method:"get" tags:"提现申请" summary:"提现申请列表"`
	*dto_withdraw.Query
}
type GetListRes struct {
	Total int                  `json:"total" dc:"总数"`
	List  []*dao_withdraw.List `json:"list" dc:"提现申请列表"`
}

type DeleteReq struct {
	g.Meta `path:"/withdraw/delete" method:"post" tags:"提现申请" summary:"删除提现申请"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}

type ApplyReq struct {
	g.Meta `path:"/withdraw/apply" method:"post" tags:"提现申请" summary:"提现佣金"`
	*dto_withdraw.Apply
}
type ApplyRes struct{}
