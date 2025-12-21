// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package upload

import (
	"context"

	"server/app/admin/api/upload/v1"
)

type IUploadV1 interface {
	UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error)
	GetChunkIdentifier(ctx context.Context, req *v1.GetChunkIdentifierReq) (res *v1.GetChunkIdentifierRes, err error)
	UploaChunk(ctx context.Context, req *v1.UploaChunkReq) (res *v1.UploaChunkRes, err error)
	UploaChunkMerge(ctx context.Context, req *v1.UploaChunkMergeReq) (res *v1.UploaChunkMergeRes, err error)
}
