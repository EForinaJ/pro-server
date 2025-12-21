package v1

import (
	dao_recharge "server/app/frontend/dao/recharge"
	dto_recharge "server/app/frontend/dto/recharge"

	"github.com/gogf/gf/v2/frame/g"
)

type WechatMiniProgramReq struct {
	g.Meta `path:"/recharge/wechat/mini/program" method:"post" tags:"充值" summary:"微信小程序充值订单"`
	*dto_recharge.WechatMiniProgram
}
type WechatMiniProgramRes struct {
	*dao_recharge.WechatMiniProgram
}

type GetNotifyReq struct {
	g.Meta `path:"/recharge/notify" method:"post" tags:"充值" summary:"微信支付充值回调通知"`
}
type GetNotifyRes struct{}

type GetStatusReq struct {
	g.Meta `path:"/recharge/status" method:"get" tags:"充值" summary:"获取充值订单状态"`
	Id     int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
type GetStatusRes struct {
	Status int `json:"status" dc:"支付状态码"`
}
