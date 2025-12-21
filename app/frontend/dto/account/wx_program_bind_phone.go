package dto_account

type WxProgramBindPhone struct {
	Code string `v:"required#请输入登录码" p:"code" dc:"登录码"`
}
