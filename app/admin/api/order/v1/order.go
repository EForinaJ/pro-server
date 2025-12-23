package v1

import (
	dao_order "server/app/admin/dao/order"
	dto_order "server/app/admin/dto/order"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"订单" summary:"订单列表"`
	*dto_order.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"总数"`
	List  []*dao_order.List `json:"list" dc:"订单列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" tags:"订单" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_order.Detail
}

type RefundReq struct {
	g.Meta `path:"/order/refund" method:"post" tags:"订单" summary:"订单退款"`
	*dto_order.Refund
}
type RefundRes struct{}

type AddDiscountReq struct {
	g.Meta `path:"/order/add/discount" method:"post" tags:"订单" summary:"订单添加折扣"`
	*dto_order.AddDiscount
}
type AddDiscountRes struct{}

type DeleteReq struct {
	g.Meta `path:"/order/delete" method:"post" tags:"订单" summary:"删除订单"`
	Ids    []int64 `json:"ids" v:"required|array#ids不能为空|删除列表是一个数组"`
}
type DeleteRes struct{}

type PaidReq struct {
	g.Meta `path:"/order/paid" method:"post" tags:"订单" summary:"订单确认收款"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type PaidRes struct{}

type CancelReq struct {
	g.Meta `path:"/order/cancel" method:"post" tags:"订单" summary:"关闭订单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type CancelRes struct{}
