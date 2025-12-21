package dto_auth

type ProgramLogin struct {
	Provider string `v:"required#请输入登录服务商" p:"provider" dc:"登录服务商"`
	Code     string `v:"required#请输入用户登录码" p:"code" dc:"用户登录码"`
}
