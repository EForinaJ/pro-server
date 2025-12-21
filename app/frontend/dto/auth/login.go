package dto_auth

type Login struct {
	Phone    string `v:"required#请输入用户手机号" p:"phone" dc:"用户手机号"`
	PassWord string `v:"required#请输入用户登录码" p:"password" dc:"用户密码"`
}
