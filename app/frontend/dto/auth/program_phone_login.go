package dto_auth

type WxProgramPhoneLogin struct {
	Code string `v:"required#请输入登录码" p:"code" dc:"登录码"`
}
