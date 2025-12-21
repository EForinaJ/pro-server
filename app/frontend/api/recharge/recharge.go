// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package recharge

import (
	"context"

	"server/app/frontend/api/recharge/v1"
)

type IRechargeV1 interface {
	WechatMiniProgram(ctx context.Context, req *v1.WechatMiniProgramReq) (res *v1.WechatMiniProgramRes, err error)
	GetNotify(ctx context.Context, req *v1.GetNotifyReq) (res *v1.GetNotifyRes, err error)
	GetStatus(ctx context.Context, req *v1.GetStatusReq) (res *v1.GetStatusRes, err error)
}
