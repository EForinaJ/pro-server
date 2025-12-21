package dao_order

import "github.com/gogf/gf/v2/os/gtime"

type Detail struct {
	Id             int64         `json:"id" dc:"ID"`
	Code           string        `json:"code" dc:"订单号"`
	User           *User         `json:"user" dc:"下单用户"`
	Product        *Product      `json:"product" dc:"订单商品"`
	TotalAmount    float64       `json:"totalAmount" dc:"订单总额"`
	DiscountAmount float64       `json:"discountAmount" dc:"订单优惠"`
	ActualAmount   float64       `json:"actualAmount" dc:"实付金额"`
	PayAmount      float64       `json:"payAmount" dc:"已支付"`
	Requirements   string        `json:"requirements" dc:"订单要求"`
	PayMode        int           `json:"payMode" dc:"收款方式"`
	PayStatus      int           `json:"payStatus" dc:"支付状态"`
	PayTime        *gtime.Time   `json:"payTime" dc:"支付时间"`
	CreateTime     *gtime.Time   `json:"createTime" dc:"下单时间"`
	StartTime      *gtime.Time   `json:"startTime" dc:"开始时间"`
	FinishTime     *gtime.Time   `json:"finishTime" dc:"结束时间"`
	Status         int           `json:"status" dc:"订单状态"`
	Remark         string        `json:"remark" dc:"管理备注"`
	WitkeyList     []*WitkeyList `json:"witkeyList" dc:"派发威客"`
}
