package v1

import (
	dto_upload "server/app/admin/dto/upload"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `path:"/upload/file" method:"post" tags:"上传" summary:"小文件上传"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}
type UploadFileRes struct {
	Links []string `json:"links" dc:"返回上传的链接"`
}

type GetChunkIdentifierReq struct {
	g.Meta     `path:"/upload/chunkIdentifier" method:"get" tags:"上传" summary:"获取已上传的切片列表"`
	Identifier string `json:"identifier"  dc:"切片标识符"`
}
type GetChunkIdentifierRes struct {
	List []int `json:"result" dc:"返回切片的标识列表"`
}

type UploaChunkReq struct {
	g.Meta `path:"/upload/chunk" method:"post" tags:"上传" summary:"上传切片"`
	*dto_upload.Chunk
}
type UploaChunkRes struct{}

type UploaChunkMergeReq struct {
	g.Meta `path:"/upload/chunk/merge" method:"post" tags:"上传" summary:"合并切片"`
	*dto_upload.Merge
}
type UploaChunkMergeRes struct {
	Links []string `json:"links" dc:"返回上传的链接"`
}
