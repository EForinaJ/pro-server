package dto_auth

type Login struct {
	Phone    string `v:"required#请输入手机号" p:"phone" dc:"用户手机号"`
	PassWord string `v:"required#请输入密码" p:"password" dc:"用户登录密码"`
}
