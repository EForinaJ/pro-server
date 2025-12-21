package dto_account

type CreateCard struct {
	Type     int    `p:"type" v:"required#请选择类型" dc:"类型"`
	Name     string `p:"name" v:"required#请输入账户名称" dc:"名称"`
	NickName string `p:"nickName" v:"required#请输入账户昵称" dc:"昵称"`
	Number   string `p:"number" v:"required#请输入账户号码" dc:"号码"`
}
