package dto_order

type Log struct {
	Page  int   `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit int   `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Id    int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
