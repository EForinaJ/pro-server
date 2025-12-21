package dto_project

type Query struct {
	Page   int    `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit  int    `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Code   string `p:"code" dc:"派单号"`
	Order  string `p:"order" dc:"订单号"`
	Status int    `p:"status" dc:"派单状态"`
}
