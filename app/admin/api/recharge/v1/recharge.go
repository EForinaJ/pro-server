package v1

import (
	dao_recharge "server/app/admin/dao/recharge"
	dto_recharge "server/app/admin/dto/recharge"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/recharge/list" method:"get" tags:"充值" summary:"充值记录列表"`
	*dto_recharge.Query
}
type GetListRes struct {
	Total int                  `json:"total" dc:"总数"`
	List  []*dao_recharge.List `json:"list" dc:"充值记录列表"`
}

type RevokeRechargeReq struct {
	g.Meta `path:"/recharge/revoke" method:"post" tags:"充值" summary:"撤回预存充值"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type RevokeRechargeRes struct{}
