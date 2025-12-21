package dto_upload

import "github.com/gogf/gf/v2/net/ghttp"

type Chunk struct {
	File        *ghttp.UploadFile `p:"file" type:"file" dc:"选择上传文件"`
	Identifier  string            `p:"identifier"  v:"required#请上传文件唯一标识" dc:"切片标识"`
	ChunkNumber string            `p:"chunkNumber"  v:"required#请设置切片索引" dc:"切片索引"`
	TotalChunks string            `p:"totalChunks"  v:"required#请设置切片总数" dc:"切片总数"`
}
