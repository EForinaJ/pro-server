package upload

import (
	"context"

	v1 "server/app/admin/api/upload/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) UploaChunk(ctx context.Context, req *v1.UploaChunkReq) (res *v1.UploaChunkRes, err error) {
	err = service.Upload().UploadChunk(ctx, req.Chunk)
	if err != nil {
		return nil, err
	}
	return
}
