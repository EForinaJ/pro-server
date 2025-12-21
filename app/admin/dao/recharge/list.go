package dao_recharge

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Code       string      `json:"code" dc:"充值号"`
	Name       string      `json:"name" dc:"昵称"`
	Money      float64     `json:"money" dc:"提现金额"`
	Remark     string      `json:"remark" dc:"备注"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
