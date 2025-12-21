package dto_user

type Query struct {
	Page   int    `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit  int    `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Id     int64  `p:"id" dc:"用户id"`
	Name   string `p:"name" dc:"用户名称"`
	Phone  string `p:"phone" dc:"手机号"`
	Status int    `p:"status" dc:"状态"`
}

type RechargeQuery struct {
	Page  int   `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit int   `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Id    int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}

type BalanceQuery struct {
	Page  int   `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit int   `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Id    int64 `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
}
