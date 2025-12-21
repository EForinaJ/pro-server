package v1

import (
	dao_order "server/app/frontend/dao/order"
	dto_order "server/app/frontend/dto/order"

	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"订单" summary:"获取订单列表"`
	*dto_order.Query
}
type GetListRes struct {
	Total int               `json:"total" dc:"总数"`
	List  []*dao_order.List `json:"list" dc:"列表"`
}

type GetDetailReq struct {
	g.Meta `path:"/order/detail" method:"get" tags:"订单" summary:"获取信息"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetDetailRes struct {
	*dao_order.Detail
}

type CancelReq struct {
	g.Meta `path:"/order/cancel" method:"post" tags:"订单" summary:"取消订单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type CancelRes struct{}

type DeleteReq struct {
	g.Meta `path:"/order/delete" method:"post" tags:"订单" summary:"删除订单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type DeleteRes struct{}

type BalancePayReq struct {
	g.Meta `path:"/order/balance/pay" method:"post" tags:"订单" summary:"余额支付"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type BalancePayRes struct{}

type WechatMiniProgramPayReq struct {
	g.Meta `path:"/order/wechat/mini/program/pay" method:"post" tags:"订单" summary:"小程序支付订单"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type WechatMiniProgramPayRes struct {
	*dao_order.WechatMiniProgram
}

type GetNotifyReq struct {
	g.Meta `path:"/order/notify" method:"post" tags:"充值" summary:"微信支付充值回调通知"`
}
type GetNotifyRes struct{}

type GetStatusReq struct {
	g.Meta `path:"/order/status" method:"get" tags:"订单" summary:"获取订单状态"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetStatusRes struct {
	Status int `json:"status" dc:"支付状态码"`
}
