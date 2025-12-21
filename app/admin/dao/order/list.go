package dao_order

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id         int64       `json:"id" dc:"ID"`
	Code       string      `json:"code" dc:"订单号"`
	User       *User       `json:"user" dc:"下单用户"`
	Product    *Product    `json:"product" dc:"订单商品"`
	PayAmount  float64     `json:"payAmount" dc:"订单总额"`
	Status     int         `json:"status" dc:"订单状态"`
	PayStatus  int         `json:"payStatus" dc:"支付状态"`
	PayMode    int         `json:"payMode" dc:"支付方式"`
	CreateTime *gtime.Time `json:"createTime" dc:"下单时间"`
}

type WitkeyList struct {
	Id         int64       `json:"id" dc:"项目Id"`
	Name       string      `json:"name" dc:"威客名称"`
	Game       string      `json:"game" dc:"游戏领域"`
	Title      string      `json:"title" dc:"头衔勋章"`
	Status     int         `json:"status" dc:"订单状态"`
	CreateTime *gtime.Time `json:"createTime" dc:"下单时间"`
}
