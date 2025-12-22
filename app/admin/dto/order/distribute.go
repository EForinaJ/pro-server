package dto_order

type Distribute struct {
	Id       int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	WitkeyId int64 `p:"WitkeyId" v:"required#请设置打手陪陪" dc:"打手陪陪ID"`
}
