package dto_aftersales

type Create struct {
	Type    int      `p:"type" v:"required|in:1,2,3#请输入状态|状态值只能在1或2或3"    dc:"售后类型"`
	Content string   `p:"content" v:"required#请输入拒绝售后原因" dc:"拒绝售后原因"`
	Images  []string `p:"images" dc:"售后凭证"`
	OrderId int64    `p:"orderId" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"订单id"`
}
