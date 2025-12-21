package dao_order

import "github.com/gogf/gf/v2/os/gtime"

type List struct {
	Id             int64       `json:"id" dc:"ID"`
	Product        *Product    `json:"product" dc:"订单商品"`
	Specifications string      `json:"specifications" dc:"商品规格"`
	Quantity       int         `json:"quantity" dc:"数量"`
	UnitPrice      float64     `json:"unitPrice" dc:"单价"`
	PayAmount      float64     `json:"payAmount" dc:"支付金额"`
	Status         int         `json:"status" dc:"订单状态"`
	CreateTime     *gtime.Time `json:"createTime" dc:"下单时间"`
}
