package v1

import (
	dao_game "server/app/admin/dao/game"
	dto_game "server/app/admin/dto/game"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllReq struct {
	g.Meta `path:"/game/all" method:"get" tags:"游戏" summary:"所以游戏"`
}
type GetAllRes struct {
	List []*dao_game.List `json:"list" dc:"游戏列表"`
}

type GetListReq struct {
	g.Meta `path:"/game/list" method:"get" tags:"游戏" summary:"游戏列表"`
	*dto_game.Query
}
type GetListRes struct {
	Total int              `json:"total" dc:"总数"`
	List  []*dao_game.List `json:"list" dc:"游戏列表"`
}

type CreateReq struct {
	g.Meta `path:"/game/create" method:"post" tags:"游戏" summary:"创建游戏"`
	*dto_game.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/game/edit" method:"get" tags:"游戏" summary:"获取编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetEditRes struct {
	*dao_game.Edit
}

type EditReq struct {
	g.Meta `path:"/game/edit" method:"post" tags:"游戏" summary:"游戏编辑"`
	*dto_game.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/game/delete" method:"post" tags:"游戏" summary:"删除游戏"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
