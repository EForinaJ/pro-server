package dao_order

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id             int64       `json:"id" dc:"ID"`
	Code           string      `json:"code" dc:"订单号"`
	Product        *Product    `json:"product" dc:"订单商品"`
	Specifications string      `json:"specifications" dc:"商品规格"`
	Quantity       int         `json:"quantity" dc:"数量"`
	UnitPrice      float64     `json:"unitPrice" dc:"单价"`
	TotalAmount    float64     `json:"totalAmount" dc:"订单总额"`
	DiscountAmount float64     `json:"discountAmount" dc:"订单优惠"`
	ActualAmount   float64     `json:"actualAmount" dc:"实付金额"`
	PayAmount      float64     `json:"payAmount" dc:"支付金额"`
	Requirements   string      `json:"requirements" dc:"订单要求"`
	Status         int         `json:"status" dc:"订单状态"`
	CreateTime     *gtime.Time `json:"createTime" dc:"下单时间"`
	StartTime      *gtime.Time `json:"startTime" dc:"开始时间"`
	FinishTime     *gtime.Time `json:"finishTime" dc:"结束时间"`
	PayTime        *gtime.Time `json:"payTime" dc:"支付时间"`
}
