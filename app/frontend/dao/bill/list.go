package dao_bill

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"卡包ID"`
	CreateTime *gtime.Time `json:"createTime" dc:"账单时间"`
	Type       int         `json:"type" dc:"账户类型"`
	Money      float64     `json:"money" dc:"账户金额"`
	Code       string      `json:"code" dc:"账户号"`
}
