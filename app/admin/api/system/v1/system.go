package v1

import (
	dao_system "server/app/admin/dao/system"
	dto_system "server/app/admin/dto/system"

	"github.com/gogf/gf/v2/frame/g"
)

// ---------------------- 基础配置
type GetBaseReq struct {
	g.Meta `path:"/system/base" method:"get" tags:"系统配置" summary:"获取系统基本配置"`
}
type GetBaseRes struct {
	*dao_system.Base
}
type SaveBaseReq struct {
	g.Meta `path:"/system/base" method:"post" tags:"系统配置" summary:"保存系统基础配置"`
	*dto_system.Base
}
type SaveBaseRes struct{}

// ---------------------- 联系配置
type GetContactReq struct {
	g.Meta `path:"/system/contact" method:"get" tags:"系统配置" summary:"获取联系设置"`
}
type GetContactRes struct {
	*dao_system.Contact
}
type SaveContactReq struct {
	g.Meta `path:"/system/save/contact" method:"post" tags:"系统配置" summary:"保存联系设置"`
	*dto_system.Contact
}
type SaveContactRes struct{}

// ---------------------- 邮箱配置
type SaveEmailReq struct {
	g.Meta `path:"/system/save/email" method:"post" tags:"系统配置" summary:"保存系统邮箱配置"`
	*dto_system.Email
}
type SaveEmailRes struct{}

type GetEmailReq struct {
	g.Meta `path:"/system/email" method:"get" tags:"系统配置" summary:"获取系统邮箱配置"`
}
type GetEmailRes struct {
	*dao_system.Email
}

// ---------------------- 用户配置
type GetUserReq struct {
	g.Meta `path:"/system/user" method:"get" tags:"系统配置" summary:"获取用户设置"`
}
type GetUserRes struct {
	*dao_system.User
}
type SaveUserReq struct {
	g.Meta `path:"/system/user" method:"post" tags:"系统配置" summary:"保存获取用户设置"`
	*dto_system.User
}
type SaveUserRes struct{}

// ---------------------- 支付设置
type GetWithdrawReq struct {
	g.Meta `path:"/system/withdraw" method:"get" tags:"系统配置" summary:"获取系统支付设置"`
}
type GetWithdrawRes struct {
	*dao_system.Withdraw
}
type SaveWithdrawReq struct {
	g.Meta `path:"/system/withdraw" method:"post" tags:"系统配置" summary:"保存系统支付设置"`
	*dto_system.Withdraw
}
type SaveWithdrawRes struct{}

// ---------------------- 文件配置
type GetFileReq struct {
	g.Meta `path:"/system/file" method:"get" tags:"系统配置" summary:"获取系统文件设置"`
}
type GetFileRes struct {
	*dao_system.File
}
type SaveFileReq struct {
	g.Meta `path:"/system/file" method:"post" tags:"系统配置" summary:"保存系统文件设置"`
	*dto_system.File
}
type SaveFileRes struct{}

// --------------------- 协议配置
type GetUserAgreementReq struct {
	g.Meta `path:"/system/user/agreement" method:"get" tags:"系统配置" summary:"获取用户协议"`
}
type GetUserAgreementRes struct {
	Content string `json:"content"  dc:"协议内容"`
}
type SaveUserAgreementReq struct {
	g.Meta  `path:"/system/user/agreement" method:"post" tags:"系统配置" summary:"保存用户协议"`
	Content string `json:"content" p:"content" v:"required#请设置协议内容" dc:"协议内容"`
}
type SaveUserAgreementRes struct{}

type GetPrivacyPolicyReq struct {
	g.Meta `path:"/system/privacy/policy" method:"get" tags:"系统配置" summary:"获取隐私协议"`
}
type GetPrivacyPolicyRes struct {
	Content string `json:"content"  dc:"协议内容"`
}
type SavePrivacyPolicyReq struct {
	g.Meta  `path:"/system/privacy/policy" method:"post" tags:"系统配置" summary:"保存隐私协议"`
	Content string `json:"content" p:"content" v:"required#请设置协议内容" dc:"协议内容"`
}
type SavePrivacyPolicyRes struct{}

type GetAboutUsReq struct {
	g.Meta `path:"/system/about/us" method:"get" tags:"系统配置" summary:"获取关于我们"`
}
type GetAboutUsRes struct {
	Content string `json:"content"  dc:"协议内容"`
}
type SaveAboutUsReq struct {
	g.Meta  `path:"/system/about/us" method:"post" tags:"系统配置" summary:"保存关于我们"`
	Content string `json:"content" p:"content" v:"required#请设置内容" dc:"协议内容"`
}
type SaveAboutUsRes struct{}

// ---------------------- 微信小程序配置项
type GetWechatMiniProgramReq struct {
	g.Meta `path:"/system/wechat/mini/program" method:"get" tags:"系统配置" summary:"获取微信小程序设置"`
}
type GetWechatMiniProgramRes struct {
	*dao_system.WechatMiniProgram
}
type SaveWechatMiniProgramReq struct {
	g.Meta `path:"/system/save/wechat/mini/program" method:"post" tags:"系统配置" summary:"保存微信小程序设置"`
	*dto_system.WechatMiniProgram
}
type SaveWechatMiniProgramRes struct{}

// ---------------------- 微信支付配置项
type GetWechatPayReq struct {
	g.Meta `path:"/system/wechat/pay" method:"get" tags:"系统配置" summary:"获取微信支付设置"`
}
type GetWechatPayRes struct {
	*dao_system.WechatPay
}
type SaveWechatPayReq struct {
	g.Meta `path:"/system/save/wechat/pay" method:"post" tags:"系统配置" summary:"保存微信支付设置"`
	*dto_system.WechatPay
}
type SaveWechatPayRes struct{}
