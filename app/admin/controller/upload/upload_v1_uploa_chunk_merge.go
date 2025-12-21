package upload

import (
	"context"

	v1 "server/app/admin/api/upload/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) UploaChunkMerge(ctx context.Context, req *v1.UploaChunkMergeReq) (res *v1.UploaChunkMergeRes, err error) {
	list, err := service.Upload().MergeChunk(ctx, req.Merge)
	if err != nil {
		return nil, err
	}
	res = &v1.UploaChunkMergeRes{
		Links: list,
	}
	return
}
