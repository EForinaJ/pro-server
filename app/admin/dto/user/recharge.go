package dto_user

type Recharge struct {
	Money float64 `p:"money"  v:"required#请输入金额"   dc:"金额"`
	// Mode  int     `p:"mode" v:"required|in:1,2#请选择支付方式|支付方式值在1或2"   dc:"支付方式"`
	Id int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
