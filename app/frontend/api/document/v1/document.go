package v1

import (
	dao_document "server/app/frontend/dao/document"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDetailReq struct {
	g.Meta `path:"/document/detail" method:"get" tags:"文档" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"文档id"`
}
type GetDetailRes struct {
	*dao_document.Detail
}
