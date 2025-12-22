package dto_distribute

type Query struct {
	Page  int    `p:"page"  v:"required|min:1#请输入页数|值最小是1" dc:"列表页数"`
	Limit int    `p:"limit" v:"required|min:8#请输入条数|值最小是8" dc:"列表条数"`
	Name  string `json:"name"`
	Code  string `json:"code"`
}
