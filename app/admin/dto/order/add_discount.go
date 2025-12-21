package dto_order

type AddDiscount struct {
	Money float64 `p:"money" v:"required#请输入优惠金额"   dc:"优惠金额"`
	Id    int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
