// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package upload

import (
	"context"

	"server/app/frontend/api/upload/v1"
)

type IUploadV1 interface {
	UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error)
}
