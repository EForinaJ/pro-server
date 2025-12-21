package dao_aftersales

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Code       string      `json:"code" dc:"售后号"`
	Order      string      `json:"order" dc:"订单号"`
	Product    *Product    `json:"product" dc:"订单商品"`
	Type       int         `json:"type" dc:"售后类型"`
	Status     int         `json:"status" dc:"售后状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"申请时间"`
}
