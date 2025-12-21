package v1

import (
	dao_user "server/app/admin/dao/user"
	dto_user "server/app/admin/dto/user"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"用户" summary:"用户列表"`
	*dto_user.Query
}
type GetListRes struct {
	Total int              `json:"total" dc:"总数"`
	List  []*dao_user.List `json:"list" dc:"用户列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/user/detail" method:"get" tags:"用户" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_user.Detail
}

type CreateReq struct {
	g.Meta `path:"/user/create" method:"post" tags:"用户" summary:"创建用户"`
	*dto_user.Create
}
type CreateRes struct{}

type GetEidtReq struct {
	g.Meta `path:"/user/edit" method:"get" tags:"用户" summary:"获取编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetEidtRes struct {
	*dao_user.Edit
}
type EditReq struct {
	g.Meta `path:"/user/edit" method:"post" tags:"用户" summary:"修改用户"`
	*dto_user.Edit
}
type EditRes struct{}

type ChangeBalanceReq struct {
	g.Meta `path:"/user/change/balance" method:"post" tags:"用户" summary:"变更余额"`
	*dto_user.ChangeBalance
}
type ChangeBalanceRes struct{}

type RechargeReq struct {
	g.Meta `path:"/user/recharge" method:"post" tags:"用户" summary:"预存充值"`
	*dto_user.Recharge
}
type RechargeRes struct{}

type GetRechargeListReq struct {
	g.Meta `path:"/user/recharge/list" method:"get" tags:"用户" summary:"充值列表"`
	*dto_user.RechargeQuery
}
type GetRechargeListRes struct {
	Total int                      `json:"total" dc:"总数"`
	List  []*dao_user.RechargeList `json:"list" dc:"充值列表"`
}

type GetBalanceListReq struct {
	g.Meta `path:"/user/balance/list" method:"get" tags:"用户" summary:"余额日志"`
	*dto_user.BalanceQuery
}
type GetBalanceListRes struct {
	Total int                     `json:"total" dc:"总数"`
	List  []*dao_user.BalanceList `json:"list" dc:"余额日志列表"`
}

type DeleteReq struct {
	g.Meta `path:"/user/delete" method:"post" tags:"用户" summary:"删除用户"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}

type GetLogsReq struct {
	g.Meta `path:"/user/logs" method:"get" tags:"派单" summary:"派单操作日志"`
	*dto_user.Log
}
type GetLogsRes struct {
	Total int             `json:"total" dc:"总数"`
	List  []*dao_user.Log `json:"list" dc:"派单操作日志列表"`
}
