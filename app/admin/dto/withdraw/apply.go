package dto_withdraw

type Apply struct {
	Id     int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Status int    `p:"status" v:"required|in:2,3#请选择提现类型|状态值只能在2或3" dc:"提现类型"`
	Reason string `p:"reason" v:"required-if:status,3#请输入拒绝原因" dc:"拒绝原因"`
}
