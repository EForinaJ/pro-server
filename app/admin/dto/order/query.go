package dto_order

type Query struct {
	Page   int    `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit  int    `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Name   string `p:"name" dc:"用户昵称"`
	Email  string `p:"email" dc:"用户邮箱"`
	Code   string `p:"code" dc:"订单号"`
	Status int    `p:"status" dc:"订单状态"`
}
