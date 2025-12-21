// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package category

import (
	"context"

	"server/app/frontend/api/category/v1"
)

type ICategoryV1 interface {
	GetAll(ctx context.Context, req *v1.GetAllReq) (res *v1.GetAllRes, err error)
}
