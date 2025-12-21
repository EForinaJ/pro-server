package v1

import (
	dao_experience "server/app/frontend/dao/experience"
	dto_experience "server/app/frontend/dto/experience"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/experience/list" method:"get" tags:"经验" summary:"获取经验列表"`
	*dto_experience.Query
}
type GetListRes struct {
	Total int                    `json:"total" dc:"总数"`
	List  []*dao_experience.List `json:"list" dc:"列表"`
}
