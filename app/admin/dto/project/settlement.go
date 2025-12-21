package dto_project

type Settlement struct {
	Id    int64   `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Money float64 `p:"money" dc:"结算佣金"`
}
