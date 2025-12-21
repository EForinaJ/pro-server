package v1

import (
	dao_witkey "server/app/admin/dao/witkey"
	dto_witkey "server/app/admin/dto/witkey"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/witkey/list" method:"get" tags:"威客" summary:"威客列表"`
	*dto_witkey.Query
}
type GetListRes struct {
	Total int                `json:"total" dc:"总数"`
	List  []*dao_witkey.List `json:"list" dc:"威客列表"`
}

type CreateReq struct {
	g.Meta `path:"/witkey/create" method:"post" tags:"威客" summary:"创建威客"`
	*dto_witkey.Create
}
type CreateRes struct{}

type GetEidtReq struct {
	g.Meta `path:"/witkey/edit" method:"get" tags:"威客" summary:"获取编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetEidtRes struct {
	*dao_witkey.Edit
}
type EditReq struct {
	g.Meta `path:"/witkey/edit" method:"post" tags:"威客" summary:"修改威客"`
	*dto_witkey.Edit
}
type EditRes struct{}

type GetDetailReq struct {
	g.Meta `path:"/witkey/detail" method:"get" tags:"威客" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_witkey.Detail
}

type ChangeCommissionReq struct {
	g.Meta `path:"/witkey/change/commission" method:"post" tags:"威客" summary:"变更佣金"`
	*dto_witkey.Commission
}
type ChangeCommissionRes struct{}

type GetCommissionListReq struct {
	g.Meta `path:"/witkey/commission/list" method:"get" tags:"用户" summary:"余额日志"`
	*dto_witkey.CommissionQuery
}
type GetCommissionListRes struct {
	Total int                          `json:"total" dc:"总数"`
	List  []*dao_witkey.CommissionList `json:"list" dc:"余额日志列表"`
}

type PunishReq struct {
	g.Meta `path:"/witkey/punish" method:"post" tags:"威客" summary:"添加处罚"`
	*dto_witkey.Punish
}
type PunishRes struct{}

type GetPunishListReq struct {
	g.Meta `path:"/witkey/punish/list" method:"get" tags:"威客" summary:"处罚列表"`
	*dto_witkey.PunishQuery
}
type GetPunishListRes struct {
	Total int                      `json:"total" dc:"总数"`
	List  []*dao_witkey.PunishList `json:"list" dc:"处罚列表"`
}

type RevokePunishReq struct {
	g.Meta `path:"/witkey/punish/revoke" method:"post" tags:"威客" summary:"撤回处罚"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type RevokePunishRes struct{}

type DeleteReq struct {
	g.Meta `path:"/witkey/delete" method:"post" tags:"威客" summary:"删除威客"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
