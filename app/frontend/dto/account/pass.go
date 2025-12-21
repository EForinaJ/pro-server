package dto_account

type Password struct {
	OldPass string `p:"oldPass" v:"required#请输入旧密码" dc:"旧密码"`
	NewPass string `p:"newPass" v:"required#请输入新密码" dc:"新密码"`
}
