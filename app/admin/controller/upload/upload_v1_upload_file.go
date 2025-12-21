package upload

import (
	"context"

	v1 "server/app/admin/api/upload/v1"
	"server/app/admin/service"
	utils_error "server/app/common/utils/error"
	"server/app/common/utils/response"
)

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	if req.File == nil {
		return nil, utils_error.Err(response.UPDATE_FAILED)
	}

	err = service.Upload().CheckUpload(ctx, req.File)
	if err != nil {
		return nil, err
	}

	links, err := service.Upload().UploadMinFile(ctx, req.File)
	if err != nil {
		return nil, err
	}
	res = &v1.UploadFileRes{
		Links: links,
	}
	return
}
