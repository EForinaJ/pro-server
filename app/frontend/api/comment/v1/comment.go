package v1

import (
	dto_comment "server/app/frontend/dto/comment"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"/comment/create" method:"post" tags:"评价" summary:"创建评价"`
	*dto_comment.Create
}
type CreateRes struct{}
