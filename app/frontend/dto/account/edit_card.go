package dto_account

type EditCard struct {
	Id     int64  `p:"id" v:"required|integer|min:1#请输入id|id类型必须是整型|id最小为1" dc:"id"`
	Name   string `p:"name" v:"required#请输入账户名称" dc:"名称"`
	Number string `p:"number" v:"required#请输入账户号码" dc:"号码"`
}
