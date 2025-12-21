package v1

import (
	dao_project "server/app/admin/dao/project"
	dto_project "server/app/admin/dto/project"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/project/list" method:"get" tags:"项目" summary:"项目列表"`
	*dto_project.Query
}
type GetListRes struct {
	Total int                 `json:"total" dc:"总数"`
	List  []*dao_project.List `json:"list" dc:"项目列表"`
}

type GetLogsReq struct {
	g.Meta `path:"/project/logs" method:"get" tags:"项目" summary:"项目操作日志"`
	*dto_project.Log
}
type GetLogsRes struct {
	Total int                `json:"total" dc:"总数"`
	List  []*dao_project.Log `json:"list" dc:"项目操作日志列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/project/detail" method:"get" tags:"项目" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_project.Detail
}

type SettlementReq struct {
	g.Meta `path:"/project/settlement" method:"post" tags:"项目" summary:"项目结算"`
	*dto_project.Settlement
}
type SettlementRes struct{}

type CancelReq struct {
	g.Meta `path:"/project/cancel" method:"post" tags:"项目" summary:"项目取消"`
	*dto_project.Cancel
}
type CancelRes struct{}

type DeleteReq struct {
	g.Meta `path:"/project/delete" method:"post" tags:"项目" summary:"删除项目"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
