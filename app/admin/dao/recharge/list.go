package dao_recharge

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Code       string      `json:"code" dc:"充值号"`
	User       string      `json:"user" dc:"充值用户"`
	PayMode    int         `json:"payMode" dc:"充值方式"`
	Amount     float64     `json:"amount" dc:"充值金额"`
	Remark     string      `json:"remark" dc:"备注"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
