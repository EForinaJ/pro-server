// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package onboarding

import (
	"context"

	"server/app/admin/api/onboarding/v1"
)

type IOnboardingV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetDetail(ctx context.Context, req *v1.GetDetailReq) (res *v1.GetDetailRes, err error)
	Apply(ctx context.Context, req *v1.ApplyReq) (res *v1.ApplyRes, err error)
}
