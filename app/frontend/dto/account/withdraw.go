package dto_account

type Withdraw struct {
	Id    int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"提现账户ID"`
	Money float64 `p:"money" v:"required#请输入提现金额" dc:"提现金额"`
}
