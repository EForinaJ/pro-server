package v1

import (
	dao_aftersales "server/app/admin/dao/aftersales"
	dto_aftersales "server/app/admin/dto/aftersales"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/aftersales/list" method:"get" tags:"售后" summary:"售后列表"`
	*dto_aftersales.Query
}
type GetListRes struct {
	Total int                    `json:"total" dc:"总数"`
	List  []*dao_aftersales.List `json:"list" dc:"售后列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/aftersales/detail" method:"get" tags:"售后" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_aftersales.Detail
}
