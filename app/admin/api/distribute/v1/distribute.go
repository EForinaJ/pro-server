package v1

import (
	dto_distribute "server/app/admin/dto/distribute"

	"github.com/gogf/gf/v2/frame/g"
)

// type GetListReq struct {
// 	g.Meta `path:"/distribute/list" method:"get" tags:"游戏" summary:"游戏列表"`
// 	*dto_distribute.Query
// }
// type GetListRes struct {
// 	Total int              `json:"total" dc:"总数"`
// 	List  []*dao_distribute.List `json:"list" dc:"游戏列表"`
// }

type CreateReq struct {
	g.Meta `path:"/distribute/create" method:"post" tags:"派单" summary:"创建派单"`
	*dto_distribute.Create
}
type CreateRes struct{}
