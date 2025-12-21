package v1

import (
	dao_goods "server/app/frontend/dao/goods"
	dto_goods "server/app/frontend/dto/goods"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/goods/list" method:"get" tags:"商品" summary:"商品列表"`
	*dto_goods.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"总数"`
	List  []*dao_goods.List `json:"list" dc:"商品列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"商品" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"文档id"`
}
type GetDetailRes struct {
	*dao_goods.Detail
}
