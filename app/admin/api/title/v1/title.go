package v1

import (
	dao_title "server/app/admin/dao/title"
	dto_title "server/app/admin/dto/title"

	"github.com/gogf/gf/v2/frame/g"
)

type GetAllReq struct {
	g.Meta `path:"/title/all" method:"get" tags:"头衔" summary:"获取某个分类下的头衔"`
	GameId int64 `p:"gameId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetAllRes struct {
	List []*dao_title.OptionsList `json:"list" dc:"头衔列表"`
}

type GetListReq struct {
	g.Meta `path:"/title/list" method:"get" tags:"头衔" summary:"头衔列表"`
	*dto_title.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"总数"`
	List  []*dao_title.List `json:"list" dc:"头衔列表"`
}

type CreateReq struct {
	g.Meta `path:"/title/create" method:"post" tags:"头衔" summary:"创建头衔"`
	*dto_title.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/title/edit" method:"get" tags:"头衔" summary:"获取编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetEditRes struct {
	*dao_title.Edit
}

type EditReq struct {
	g.Meta `path:"/title/edit" method:"post" tags:"头衔" summary:"头衔编辑"`
	*dto_title.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/title/delete" method:"post" tags:"头衔" summary:"删除头衔"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
