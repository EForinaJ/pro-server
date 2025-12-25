package dto_order

type Refund struct {
	Type   int     `p:"type" v:"required|in:1,2,3#请选择售后类型|售后类型不对" dc:"售后类型"`
	Money  float64 `p:"money" v:"required#请设置金额"`
	Reason string  `p:"reason" v:"required#请输入退款原因" dc:"退款原因"`
	Id     int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
