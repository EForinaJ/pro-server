package dao_withdraw

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Code       string      `json:"code" dc:"UID"`
	Manage     string      `json:"manage" dc:"管理昵称"`
	Witkey     string      `json:"witkey" dc:"威客昵称"`
	Money      float64     `json:"money" dc:"提现金额"`
	Remark     string      `json:"remark" dc:"备注"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"创建时间"`
}
