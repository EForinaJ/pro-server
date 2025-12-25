package dto_settlement

type Apply struct {
	Id     int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"分类id"`
	Status int    `p:"status" v:"required#请选择结算类型" dc:"结算类型"`
	Reason string `p:"reason" v:"required-if:status,3#请输入拒绝原因" dc:"拒绝原因"`
}
