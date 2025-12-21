package dto_witkey

type Punish struct {
	Money   float64 `p:"money"  dc:"处罚金额"`
	Content string  `p:"content" v:"required#请输入处罚的内容"   dc:"处罚内容"`
	Id      int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
