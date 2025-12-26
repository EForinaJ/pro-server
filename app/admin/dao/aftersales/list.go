package dao_aftersales

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"Id"`
	Code       string      `json:"code" dc:"售后编号"`
	OrderCode  string      `json:"orderCode" dc:"订单编号"`
	Product    *Product    `json:"product" dc:"商品信息"`
	User       string      `json:"user" dc:"申请用户"`
	Amount     float64     `json:"amount" dc:"退款金额"`
	Status     int         `json:"status" dc:"状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"下单时间"`
}

type Product struct {
	Id   int64  `json:"id" dc:"Id"`
	Pic  string `json:"pic" dc:"商品图片"`
	Name string `json:"name" dc:"订单编号"`
}
