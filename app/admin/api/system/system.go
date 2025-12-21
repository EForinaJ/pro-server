// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"server/app/admin/api/system/v1"
)

type ISystemV1 interface {
	GetBase(ctx context.Context, req *v1.GetBaseReq) (res *v1.GetBaseRes, err error)
	SaveBase(ctx context.Context, req *v1.SaveBaseReq) (res *v1.SaveBaseRes, err error)
	GetContact(ctx context.Context, req *v1.GetContactReq) (res *v1.GetContactRes, err error)
	SaveContact(ctx context.Context, req *v1.SaveContactReq) (res *v1.SaveContactRes, err error)
	SaveEmail(ctx context.Context, req *v1.SaveEmailReq) (res *v1.SaveEmailRes, err error)
	GetEmail(ctx context.Context, req *v1.GetEmailReq) (res *v1.GetEmailRes, err error)
	GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error)
	SaveUser(ctx context.Context, req *v1.SaveUserReq) (res *v1.SaveUserRes, err error)
	GetWithdraw(ctx context.Context, req *v1.GetWithdrawReq) (res *v1.GetWithdrawRes, err error)
	SaveWithdraw(ctx context.Context, req *v1.SaveWithdrawReq) (res *v1.SaveWithdrawRes, err error)
	GetFile(ctx context.Context, req *v1.GetFileReq) (res *v1.GetFileRes, err error)
	SaveFile(ctx context.Context, req *v1.SaveFileReq) (res *v1.SaveFileRes, err error)
	GetUserAgreement(ctx context.Context, req *v1.GetUserAgreementReq) (res *v1.GetUserAgreementRes, err error)
	SaveUserAgreement(ctx context.Context, req *v1.SaveUserAgreementReq) (res *v1.SaveUserAgreementRes, err error)
	GetPrivacyPolicy(ctx context.Context, req *v1.GetPrivacyPolicyReq) (res *v1.GetPrivacyPolicyRes, err error)
	SavePrivacyPolicy(ctx context.Context, req *v1.SavePrivacyPolicyReq) (res *v1.SavePrivacyPolicyRes, err error)
	GetAboutUs(ctx context.Context, req *v1.GetAboutUsReq) (res *v1.GetAboutUsRes, err error)
	SaveAboutUs(ctx context.Context, req *v1.SaveAboutUsReq) (res *v1.SaveAboutUsRes, err error)
	GetWechatMiniProgram(ctx context.Context, req *v1.GetWechatMiniProgramReq) (res *v1.GetWechatMiniProgramRes, err error)
	SaveWechatMiniProgram(ctx context.Context, req *v1.SaveWechatMiniProgramReq) (res *v1.SaveWechatMiniProgramRes, err error)
	GetWechatPay(ctx context.Context, req *v1.GetWechatPayReq) (res *v1.GetWechatPayRes, err error)
	SaveWechatPay(ctx context.Context, req *v1.SaveWechatPayReq) (res *v1.SaveWechatPayRes, err error)
}
