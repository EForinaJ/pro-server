package dao_aftersales

import (
	dao_order "server/app/admin/dao/order"

	"github.com/gogf/gf/v2/os/gtime"
)

type Detail struct {
	Id         int64             `json:"id" dc:"Id"`
	Code       string            `json:"code" dc:"售后编号"`
	Order      *dao_order.Detail `json:"order" dc:"订单信息"`
	Manage     string            `json:"manage" dc:"审核者"`
	Amount     float64           `json:"amount" dc:"退款金额"`
	Type       int               `json:"type" dc:"退款类型"`
	Reason     string            `json:"reason" dc:"退款原因"`
	Status     int               `json:"status" dc:"状态"`
	CreateTime *gtime.Time       `json:"createTime" dc:"下单时间"`
}
