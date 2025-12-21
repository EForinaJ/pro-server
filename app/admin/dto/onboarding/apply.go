package dto_onboarding

type Apply struct {
	Id           int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Status       int    `p:"status" v:"required|in:2,3#请输入状态|状态值只能在2或3"    dc:"审核状态"`
	RefuseReason string `p:"refuseReason" v:"required-if:status,3#请输入拒绝原因"  dc:"拒绝原因"`
}
