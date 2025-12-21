package dao_project

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Code       string      `json:"code" dc:"项目号"`
	Order      string      `json:"order" dc:"订单号"`
	Witkey     *Witkey     `json:"witkey" dc:"接单者"`
	Product    *Product    `json:"product" dc:"项目商品"`
	Money      float64     `json:"money"  dc:"项目金额"` //
	Remark     string      `json:"remark" dc:"备注"`
	Status     int         `json:"status" dc:"项目状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"项目时间"`
}
