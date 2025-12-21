package upload

import (
	"context"

	v1 "server/app/admin/api/upload/v1"
	"server/app/admin/service"
)

func (c *ControllerV1) GetChunkIdentifier(ctx context.Context, req *v1.GetChunkIdentifierReq) (res *v1.GetChunkIdentifierRes, err error) {
	list, err := service.Upload().ChunkIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, err
	}
	res = &v1.GetChunkIdentifierRes{
		List: list,
	}
	return
}
