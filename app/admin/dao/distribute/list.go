package dao_distribute

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"项目Id"`
	Code       string      `json:"code" dc:"订单编号"`
	Manage     string      `json:"manage" dc:"派单者"`
	Witkey     string      `json:"witkey" dc:"接单威客"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	IsCancel   int         `json:"isCancel" dc:"是否取消"`
	Reason     string      `json:"reason" dc:"变更原因"`
	CreateTime *gtime.Time `json:"createTime" dc:"下单时间"`
}
