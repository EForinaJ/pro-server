package dao_capital

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Related    string      `json:"related" dc:"关联单号"`
	Code       string      `json:"code" dc:"支付号"`
	Type       int         `json:"type" dc:"类型"`
	User       string      `json:"user" dc:"支付用户"`
	Mode       int         `json:"mode" dc:"支付方式"`
	Amount     float64     `json:"amount" dc:"金额"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
