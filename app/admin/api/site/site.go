// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package site

import (
	"context"

	"server/app/admin/api/site/v1"
)

type ISiteV1 interface {
	GetSite(ctx context.Context, req *v1.GetSiteReq) (res *v1.GetSiteRes, err error)
	GetUserOptions(ctx context.Context, req *v1.GetUserOptionsReq) (res *v1.GetUserOptionsRes, err error)
	GetGameOptions(ctx context.Context, req *v1.GetGameOptionsReq) (res *v1.GetGameOptionsRes, err error)
	GetCategoryeOptions(ctx context.Context, req *v1.GetCategoryeOptionsReq) (res *v1.GetCategoryeOptionsRes, err error)
	GetTitleOptions(ctx context.Context, req *v1.GetTitleOptionsReq) (res *v1.GetTitleOptionsRes, err error)
}
