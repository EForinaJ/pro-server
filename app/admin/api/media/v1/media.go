package v1

import (
	dao_media "server/app/admin/dao/media"
	dto_media "server/app/admin/dto/media"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/media/list" method:"get" tags:"媒体附件" summary:"获取媒体附件"`
	*dto_media.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"媒体附件总数"`
	List  []*dao_media.List `json:"list" dc:"媒体附件列表"`
}

type DeleteReq struct {
	g.Meta `path:"/media/delete" method:"delete" tags:"媒体附件" summary:"删除媒体附件"`
	Ids    []int64 `json:"ids" v:"required#ids不能为空" dc:"附件ID"`
}

type DeleteRes struct{}
