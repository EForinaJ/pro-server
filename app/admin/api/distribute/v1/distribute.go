package v1

import (
	dao_distribute "server/app/admin/dao/distribute"
	dto_distribute "server/app/admin/dto/distribute"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/distribute/list" method:"get" tags:"派单" summary:"派单列表"`
	*dto_distribute.Query
}
type GetListRes struct {
	Total int                    `json:"total" dc:"总数"`
	List  []*dao_distribute.List `json:"list" dc:"派单列表"`
}

type CreateReq struct {
	g.Meta `path:"/distribute/create" method:"post" tags:"派单" summary:"创建派单"`
	*dto_distribute.Create
}
type CreateRes struct{}

type CancelReq struct {
	g.Meta `path:"/distribute/cancel" method:"post" tags:"派单" summary:"取消派单"`
	*dto_distribute.Cancel
}
type CancelRes struct{}
