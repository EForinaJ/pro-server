package dto_user

type MarginDeposit struct {
	Money float64 `p:"money"  dc:"所需的变更金额"`
	Mode  int     `p:"mode" v:"required|in:1,2#请选择变更组别|变更组别值在1或2"   dc:"变更组别"`
	Id    int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
