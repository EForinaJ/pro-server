package dao_settlement

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"项目Id"`
	OrderCode  string      `json:"orderCode" dc:"订单编号"`
	Code       string      `json:"code" dc:"报单号"`
	Witkey     string      `json:"witkey" dc:"报单威客"`
	Amount     float64     `json:"amount" dc:"结算金额"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"下单时间"`
}
