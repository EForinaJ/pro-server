package v1

import (
	dao_product "server/app/admin/dao/product"
	dto_product "server/app/admin/dto/product"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/product/list" method:"get" tags:"商品" summary:"商品列表"`
	*dto_product.Query
}
type GetListRes struct {
	Total int                 `json:"total" dc:"总数"`
	List  []*dao_product.List `json:"list" dc:"商品列表"`
}

type CreateReq struct {
	g.Meta `path:"/product/create" method:"post" tags:"商品" summary:"创建商品"`
	*dto_product.Create
}
type CreateRes struct{}

type GetEditReq struct {
	g.Meta `path:"/product/edit" method:"get" tags:"商品" summary:"获取编辑信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"威客id"`
}
type GetEditRes struct {
	*dao_product.Edit
}
type EditReq struct {
	g.Meta `path:"/product/edit" method:"post" tags:"商品" summary:"商品编辑"`
	*dto_product.Edit
}
type EditRes struct{}

type DeleteReq struct {
	g.Meta `path:"/product/delete" method:"post" tags:"商品" summary:"删除商品"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}
