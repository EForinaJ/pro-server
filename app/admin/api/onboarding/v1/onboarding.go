package v1

import (
	dao_onboarding "server/app/admin/dao/onboarding"
	dto_onboarding "server/app/admin/dto/onboarding"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/onboarding/list" method:"get" tags:"入职" summary:"入职列表"`
	*dto_onboarding.Query
}
type GetListRes struct {
	Total int                    `json:"total" dc:"总数"`
	List  []*dao_onboarding.List `json:"list" dc:"入职列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/onboarding/detail" method:"get" tags:"入职" summary:"入职信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_onboarding.Detail
}

type ApplyReq struct {
	g.Meta `path:"/onboarding/apply" method:"post" tags:"入职" summary:"入职审核"`
	*dto_onboarding.Apply
}
type ApplyRes struct{}
