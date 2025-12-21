package v1

import (
	dao_level "server/app/frontend/dao/level"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllReq struct {
	g.Meta `path:"/level/all" method:"get" tags:"等级" summary:"所以等级"`
}
type GetAllRes struct {
	List []*dao_level.List `json:"list" dc:"等级列表"`
}
