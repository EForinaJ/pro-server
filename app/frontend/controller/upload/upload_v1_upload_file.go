package upload

import (
	"context"

	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
	v1 "server/app/frontend/api/upload/v1"
	"server/app/frontend/service"
)

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	if req.File == nil {
		return nil, utils_error.Err(response.UPDATE_FAILED)
	}
	links, err := service.Upload().UploadFile(ctx, req.File)
	if err != nil {
		return nil, err
	}
	res = &v1.UploadFileRes{
		Links: links,
	}
	return
}
